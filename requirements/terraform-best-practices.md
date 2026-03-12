# Welcome

[Terraform](https://www.terraform.io) is powerful (if not the most powerful out there now) and one of the most used tools which allow management of infrastructure as code. It allows developers to do a lot of things and does not restrict them from doing things in ways that will be hard to support or integrate with.

Some information described in this book may not seem like the best practices. I know this, and to help readers to separate what are established best practices and what is just another opinionated way of doing things, I sometimes use hints to provide some context and icons to specify the level of maturity on each subsection related to best practices.

The book was started in sunny Madrid in 2018, available for free here at [https://www.terraform-best-practices.com/](https://www.terraform-best-practices.com).

A few years later it has been updated with more actual best practices available with Terraform 1.0. Eventually, this book should contain most of the indisputable best practices and recommendations for Terraform users.

# Key concepts

The official Terraform documentation describes [all aspects of configuration in details](https://www.terraform.io/docs/configuration/index.html). Read it carefully to understand the rest of this section.

This section describes key concepts which are used inside the book.

## Resource

Resource is `aws_vpc`, `aws_db_instance`, etc. A resource belongs to a provider, accepts arguments, outputs attributes, and has a lifecycle. A resource can be created, retrieved, updated, and deleted.

## Resource module

Resource module is a collection of connected resources which together perform the common action (for e.g., [AWS VPC Terraform module](https://github.com/terraform-aws-modules/terraform-aws-vpc/) creates VPC, subnets, NAT gateway, etc). It depends on provider configuration, which can be defined in it, or in higher-level structures (e.g., in infrastructure module).

## Infrastructure module

An infrastructure module is a collection of resource modules, which can be logically not connected, but in the current situation/project/setup serves the same purpose. It defines the configuration for providers, which is passed to the downstream resource modules and to resources. It is normally limited to work in one entity per logical separator (e.g., AWS Region, Google Project).

For example, [terraform-aws-atlantis](https://github.com/terraform-aws-modules/terraform-aws-atlantis/) module uses resource modules like [terraform-aws-vpc](https://github.com/terraform-aws-modules/terraform-aws-vpc/) and [terraform-aws-security-group](https://github.com/terraform-aws-modules/terraform-aws-security-group/) to manage the infrastructure required for running [Atlantis](https://www.runatlantis.io) on [AWS Fargate](https://aws.amazon.com/fargate/).

Another example is [terraform-aws-cloudquery](https://github.com/cloudquery/terraform-aws-cloudquery) module where multiple modules by [terraform-aws-modules](https://github.com/terraform-aws-modules/) are being used together to manage the infrastructure as well as using Docker resources to build, push, and deploy Docker images. All in one set.

## Composition

Composition is a collection of infrastructure modules, which can span across several logically separated areas (e.g.., AWS Regions, several AWS accounts). Composition is used to describe the complete infrastructure required for the whole organization or project.

A composition consists of infrastructure modules, which consist of resources modules, which implement individual resources.

![Simple infrastructure composition](https://118479043-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2Fe1Mp2scOX6OnQbifCen3%2Fuploads%2F6DQkdvTnhqwFfVx2btjq%2Fcomposition-1.png?alt=media)

## Data source

Data source performs a read-only operation and is dependant on provider configuration, it is used in a resource module and an infrastructure module.

Data source `terraform_remote_state` acts as a glue for higher-level modules and compositions.

The [external](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) data source allows an external program to act as a data source, exposing arbitrary data for use elsewhere in the Terraform configuration. Here is an example from the [terraform-aws-lambda module](https://github.com/terraform-aws-modules/terraform-aws-lambda/blob/258e82b50adc451f51544a2b57fd1f6f8f4a61e4/package.tf#L5-L7) where the filename is computed by calling an external Python script.

The [http](https://registry.terraform.io/providers/hashicorp/http/latest/docs/data-sources/http) data source makes an HTTP GET request to the given URL and exports information about the response which is often useful to get information from endpoints where a native Terraform provider does not exist.

## Remote state

Store [Terraform state](https://www.terraform.io/docs/language/state/index.html) for each infrastructure module and composition in a remote backend, configured with ACLs, versioning, and logging. This single, authoritative source of truth keeps environments consistent and typically includes disaster-recovery features such as automated backups. Managing state locally can lead to collaboration issues and race conditions when multiple developers run Terraform at the same time, resulting in unpredictable outcomes.

## Provider, provisioner, etc

Providers, provisioners, and a few other terms are described very well in the official documentation and there is no point to repeat it here. To my opinion, they have little to do with writing good Terraform modules.

## Why so *difficult*?

While individual resources are like atoms in the infrastructure, resource modules are molecules (consisting of atoms). A module is the smallest versioned and shareable unit. It has an exact list of arguments, implement basic logic for such a unit to do the required function. e.g., [terraform-aws-security-group](https://github.com/terraform-aws-modules/terraform-aws-security-group) module creates `aws_security_group` and `aws_security_group_rule` resources based on input. This resource module by itself can be used together with other modules to create the infrastructure module.

Access to data across molecules (resource modules and infrastructure modules) is performed using the modules' outputs and data sources.

Access between compositions is often performed using remote state data sources. There are [multiple ways to share data between configurations](https://www.terraform.io/docs/language/state/remote-state-data.html#alternative-ways-to-share-data-between-configurations).

When putting concepts described above in pseudo-relations it may look like this:

```
composition-1 {
  infrastructure-module-1 {
    data-source-1 => d1

    resource-module-1 {
      data-source-2 => d2
      resource-1 (d1, d2)
      resource-2 (d2)
    }

    resource-module-2 {
      data-source-3 => d3
      resource-3 (d1, d3)
      resource-4 (d3)
    }
  }
}
```

# Code structure

Questions related to Terraform code structure are by far the most frequent in the community. Everyone thought about the best code structure for the project at some point also.

## How should I structure my Terraform configurations?

This is one of the questions where lots of solutions exist and it is very hard to give universal advice, so let's start with understanding what are we dealing with.

* What is the complexity of your project?
  * Number of related resources
  * Number of Terraform providers (see note below about "logical providers")
* How often does your infrastructure change?
  * **From** once a month/week/day
  * **To** continuously (every time when there is a new commit)
* Code change initiators? *Do you let the CI server update the repository when a new artifact is built?*
  * Only developers can push to the infrastructure repository
  * Everyone can propose a change to anything by opening a PR (including automated tasks running on the CI server)
* Which deployment platform or deployment service do you use?
  * AWS CodeDeploy, Kubernetes, or OpenShift require a slightly different approach
* How environments are grouped?
  * By environment, region, project

{% hint style="info" %}
*Logical providers* work entirely within Terraform's logic and very often don't interact with any other services, so we can think about their complexity as O(1). The most common logical providers include [random](https://registry.terraform.io/providers/hashicorp/random/latest/docs), [local](https://registry.terraform.io/providers/hashicorp/local/latest/docs), [terraform](https://www.terraform.io/docs/providers/terraform/index.html), [null](https://registry.terraform.io/providers/hashicorp/null/latest/docs), [time](https://registry.terraform.io/providers/hashicorp/time/latest).
{% endhint %}

## Getting started with the structuring of Terraform configurations

Putting all code in `main.tf` is a good idea when you are getting started or writing an example code. In all other cases you will be better having several files split logically like this:

* `main.tf` - call modules, locals, and data sources to create all resources
* `variables.tf` - contains declarations of variables used in `main.tf`
* `outputs.tf` - contains outputs from the resources created in `main.tf`
* `versions.tf` - contains version requirements for Terraform and providers

`terraform.tfvars` should not be used anywhere except [composition](https://www.terraform-best-practices.com/key-concepts#composition).

## How to think about Terraform configuration structure?

{% hint style="info" %}
Please make sure that you understand key concepts - [resource module](https://www.terraform-best-practices.com/key-concepts#resource-module), [infrastructure module](https://www.terraform-best-practices.com/key-concepts#infrastructure-module), and [composition](https://www.terraform-best-practices.com/key-concepts#composition), as they are used in the following examples.
{% endhint %}

### Common recommendations for structuring code

* It is easier and faster to work with a smaller number of resources
  * `terraform plan` and `terraform apply` both make cloud API calls to verify the status of resources
  * If you have your entire infrastructure in a single composition this can take some time
* A blast radius (in case of security breach) is smaller with fewer resources
  * Insulating unrelated resources from each other by placing them in separate compositions reduces the risk if something goes wrong
* Start your project using remote state because:
  * Your laptop is no place for your infrastructure source of truth
  * Managing a `tfstate` file in git is a nightmare
  * Later when infrastructure layers start to grow in multiple directions (number of dependencies or resources) it will be easier to keep things under control
* Practice a consistent structure and [naming](https://www.terraform-best-practices.com/naming) convention:
  * Like procedural code, Terraform code should be written for people to read first, consistency will help when changes happen six months from now
  * It is possible to move resources in Terraform state file but it may be harder to do if you have inconsistent structure and naming
* Keep resource modules as plain as possible
* Don't hardcode values that can be passed as variables or discovered using data sources
* Use data sources and `terraform_remote_state` specifically as a glue between infrastructure modules within the composition

In this book, example projects are grouped by *complexity* - from small to very-large infrastructures. This separation is not strict, so check other structures also.

### Orchestration of infrastructure modules and compositions

Having a small infrastructure means that there is a small number of dependencies and few resources. As the project grows the need to chain the execution of Terraform configurations, connecting different infrastructure modules, and passing values within a composition becomes obvious.

There are at least 5 distinct groups of orchestration solutions that developers use:

1. Terraform only. Very straightforward, developers have to know only Terraform to get the job done.
2. Terragrunt. Pure orchestration tool which can be used to orchestrate the entire infrastructure as well as handle dependencies. Terragrunt operates with infrastructure modules and compositions natively, so it reduces duplication of code.
3. In-house scripts. Often this happens as a starting point towards orchestration and before discovering Terragrunt.
4. Ansible or similar general purpose automation tool. Usually used when Terraform is adopted after Ansible, or when Ansible UI is actively used.
5. [Crossplane](https://crossplane.io) and other Kubernetes-inspired solutions. Sometimes it makes sense to utilize the Kubernetes ecosystem and employ a reconciliation loop feature to achieve the desired state of your Terraform configurations. View video [Crossplane vs Terraform](https://www.youtube.com/watch?v=ELhVbSdcqSY) for more information.

With that in mind, this book reviews the first two of these project structures, Terraform only and Terragrunt.

See examples of code structures for [Terraform](https://www.terraform-best-practices.com/examples/terraform) or [Terragrunt](https://www.terraform-best-practices.com/examples/terragrunt) in the next chapter.

# Small-size infrastructure with Terraform

Source: <https://github.com/antonbabenko/terraform-best-practices/tree/master/examples/small-terraform>

This example contains code as an example of structuring Terraform configurations for a small-size infrastructure, where no external dependencies are used.

{% hint style="success" %}

* Perfect to get started and refactor as you go
* Perfect for small resource modules
* Good for small and linear infrastructure modules (eg, [terraform-aws-atlantis](https://github.com/terraform-aws-modules/terraform-aws-atlantis))
* Good for a small number of resources (fewer than 20-30)
  {% endhint %}

{% hint style="warning" %}
Single state file for all resources can make the process of working with Terraform slow if the number of resources is growing (consider using an argument `-target` to limit the number of resources)
{% endhint %}

# Medium-size infrastructure with Terraform

Source: <https://github.com/antonbabenko/terraform-best-practices/tree/master/examples/medium-terraform>

This example contains code as an example of structuring Terraform configurations for a medium-size infrastructure which uses:

* 2 AWS accounts
* 2 separate environments (`prod` and `stage` which share nothing). Each environment lives in a separate AWS account
* Each environment uses a different version of the off-the-shelf infrastructure module (`alb`) sourced from [Terraform Registry](https://registry.terraform.io/)
* Each environment uses the same version of an internal module `modules/network` since it is sourced from a local directory.

{% hint style="success" %}

* Perfect for projects where infrastructure is logically separated (separate AWS accounts)
* Good when there is no need to modify resources shared between AWS accounts (one environment = one AWS account = one state file)
* Good when there is no need in the orchestration of changes between the environments
* Good when infrastructure resources are different per environment on purpose and can't be generalized (eg, some resources are absent in one environment or in some regions)
  {% endhint %}

{% hint style="warning" %}
As the project grows, it will be harder to keep these environments up-to-date with each other. Consider using infrastructure modules (off-the-shelf or internal) for repeatable tasks.
{% endhint %}

# Large-size infrastructure with Terraform

Source: <https://github.com/antonbabenko/terraform-best-practices/tree/master/examples/large-terraform>

This example contains code as an example of structuring Terraform configurations for a large-size infrastructure which uses:

* 2 AWS accounts
* 2 regions
* 2 separate environments (`prod` and `stage` which share nothing). Each environment lives in a separate AWS account and span resources between 2 regions
* Each environment uses a different version of the off-the-shelf infrastructure module (`alb`) sourced from [Terraform Registry](https://registry.terraform.io/)
* Each environment uses the same version of an internal module `modules/network` since it is sourced from a local directory.

{% hint style="info" %}
In a large project like described here the benefits of using Terragrunt become very visible. See [Code Structures examples with Terragrunt](https://www.terraform-best-practices.com/examples/terragrunt).
{% endhint %}

{% hint style="success" %}

* Perfect for projects where infrastructure is logically separated (separate AWS accounts)
* Good when there is no need to modify resources shared between AWS accounts (one environment = one AWS account = one state file)
* Good when there is no need for the orchestration of changes between the environments
* Good when infrastructure resources are different per environment on purpose and can't be generalized (eg, some resources are absent in one environment or in some regions)
  {% endhint %}

{% hint style="warning" %}
As the project grows, it will be harder to keep these environments up-to-date with each other. Consider using infrastructure modules (off-the-shelf or internal) for repeatable tasks.
{% endhint %}

# Naming conventions

## General conventions

{% hint style="info" %}
There should be no reason to not follow at least these conventions :)
{% endhint %}

{% hint style="info" %}
Beware that actual cloud resources often have restrictions in allowed names. Some resources, for example, can't contain dashes, some must be camel-cased. The conventions in this book refer to Terraform names themselves.
{% endhint %}

1. Use `_` (underscore) instead of `-` (dash) everywhere (in resource names, data source names, variable names, outputs, etc).
2. Prefer to use lowercase letters and numbers (even though UTF-8 is supported).

## Resource and data source arguments

1. Do not repeat resource type in resource name (not partially, nor completely):

{% hint style="success" %}

```
`resource "aws_route_table" "public" {}`
```

{% endhint %}

{% hint style="danger" %}

```
`resource "aws_route_table" "public_route_table" {}`
```

{% endhint %}

{% hint style="danger" %}

```
`resource "aws_route_table" "public_aws_route_table" {}`
```

{% endhint %}

2. Resource name should be named `this` if there is no more descriptive and general name available, or if the resource module creates a single resource of this type (eg, in [AWS VPC module](https://github.com/terraform-aws-modules/terraform-aws-vpc) there is a single resource of type `aws_nat_gateway` and multiple resources of type`aws_route_table`, so `aws_nat_gateway` should be named `this` and `aws_route_table` should have more descriptive names - like `private`, `public`, `database`).
3. Always use singular nouns for names.
4. Use `-` inside arguments values and in places where value will be exposed to a human (eg, inside DNS name of RDS instance).
5. Include argument `count` / `for_each` inside resource or data source block as the first argument at the top and separate by newline after it.
6. Include argument `tags,` if supported by resource, as the last real argument, following by `depends_on` and `lifecycle`, if necessary. All of these should be separated by a single empty line.
7. When using conditions in an argument`count` / `for_each` prefer boolean values instead of using `length` or other expressions.

## Code examples of `resource`

### Usage of `count` / `for_each`

{% hint style="success" %}
{% code title="main.tf" %}

```hcl
resource "aws_route_table" "public" {
  count = 2

  vpc_id = "vpc-12345678"
  # ... remaining arguments omitted
}

resource "aws_route_table" "private" {
  for_each = toset(["one", "two"])

  vpc_id = "vpc-12345678"
  # ... remaining arguments omitted
}
```

{% endcode %}
{% endhint %}

{% hint style="danger" %}
{% code title="main.tf" %}

```hcl
resource "aws_route_table" "public" {
  vpc_id = "vpc-12345678"
  count  = 2

  # ... remaining arguments omitted
}
```

{% endcode %}
{% endhint %}

### Placement of `tags`

{% hint style="success" %}
{% code title="main.tf" %}

```hcl
resource "aws_nat_gateway" "this" {
  count = 2

  allocation_id = "..."
  subnet_id     = "..."

  tags = {
    Name = "..."
  }

  depends_on = [aws_internet_gateway.this]

  lifecycle {
    create_before_destroy = true
  }
}
```

{% endcode %}
{% endhint %}

{% hint style="danger" %}
{% code title="main.tf" %}

```hcl
resource "aws_nat_gateway" "this" {
  count = 2

  tags = "..."

  depends_on = [aws_internet_gateway.this]

  lifecycle {
    create_before_destroy = true
  }

  allocation_id = "..."
  subnet_id     = "..."
}
```

{% endcode %}
{% endhint %}

### Conditions in `count`

{% hint style="success" %}
{% code title="outputs.tf" %}

```hcl
resource "aws_nat_gateway" "that" {    # Best
  count = var.create_public_subnets ? 1 : 0
}

resource "aws_nat_gateway" "this" {    # Good
  count = length(var.public_subnets) > 0 ? 1 : 0
}
```

{% endcode %}
{% endhint %}

## Variables

1. Don't reinvent the wheel in resource modules: use `name`, `description`, and `default` value for variables as defined in the "Argument Reference" section for the resource you are working with.
2. Support for validation in variables is rather limited (e.g. can't access other variables or do lookups if using a version before Terraform `1.9`). Plan accordingly because in many cases this feature is useless.
3. Use the plural form in a variable name when type is `list(...)` or `map(...)`.
4. Order keys in a variable block like this: `description` , `type`, `default`, `validation`.
5. Always include `description` on all variables even if you think it is obvious (you will need it in the future). Use the same wording as the upstream documentation when applicable.
6. Prefer using simple types (`number`, `string`, `list(...)`, `map(...)`, `any`) over specific type like `object()` unless you need to have strict constraints on each key.
7. Use specific types like `map(map(string))` if all elements of the map have the same type (e.g. `string`) or can be converted to it (e.g. `number` type can be converted to `string`).
8. Use type `any` to disable type validation starting from a certain depth or when multiple types should be supported.
9. Value `{}` is sometimes a map but sometimes an object. Use `tomap(...)` to make a map because there is no way to make an object.
10. Avoid double negatives: use positive variable names to prevent confusion. For example, use `encryption_enabled` instead of `encryption_disabled`.
11. For variables that should never be `null`, set `nullable = false`. This ensures that passing `null` uses the default value instead of `null`. If `null` is an acceptable value, you can omit nullable or set it to `true`.

## Outputs

Make outputs consistent and understandable outside of its scope (when a user is using a module it should be obvious what type and attribute of the value it returns).

1. The name of output should describe the property it contains and be less free-form than you would normally want.
2. Good structure for the name of output looks like `{name}_{type}_{attribute}` , where:
   1. `{name}` is a resource or data source name
      * `{name}` for `data "aws_subnet" "private"` is `private`
      * `{name}` for `resource "aws_vpc_endpoint_policy" "test"` is `test`
   2. `{type}` is a resource or data source type without a provider prefix
      * `{type}` for `data "aws_subnet" "private"` is `subnet`
      * `{type}` for `resource "aws_vpc_endpoint_policy" "test"` is `vpc_endpoint_policy`
   3. `{attribute}` is an attribute returned by the output
   4. [See examples](#code-examples-of-output).
3. If the output is returning a value with interpolation functions and multiple resources, `{name}` and `{type}` there should be as generic as possible (`this` as prefix should be omitted). [See example](#code-examples-of-output).
4. If the returned value is a list it should have a plural name. [See example](#use-plural-name-if-the-returning-value-is-a-list).
5. Always include `description` for all outputs even if you think it is obvious.
6. Avoid setting `sensitive` argument unless you fully control usage of this output in all places in all modules.
7. Prefer `try()` (available since Terraform 0.13) over `element(concat(...))` (legacy approach for the version before 0.13)

### Code examples of `output`

Return at most one ID of security group:

{% hint style="success" %}
{% code title="outputs.tf" %}

```hcl
output "security_group_id" {
  description = "The ID of the security group"
  value       = try(aws_security_group.this[0].id, aws_security_group.name_prefix[0].id, "")
}
```

{% endcode %}
{% endhint %}

When having multiple resources of the same type, `this` should be omitted in the name of output:

{% hint style="danger" %}
{% code title="outputs.tf" %}

```hcl
output "this_security_group_id" {
  description = "The ID of the security group"
  value       = element(concat(coalescelist(aws_security_group.this.*.id, aws_security_group.web.*.id), [""]), 0)
}
```

{% endcode %}
{% endhint %}

### Use plural name if the returning value is a list

{% hint style="success" %}
{% code title="outputs.tf" %}

```hcl
output "rds_cluster_instance_endpoints" {
  description = "A list of all cluster instance endpoints"
  value       = aws_rds_cluster_instance.this.*.endpoint
}
```

{% endcode %}
{% endhint %}

# Code styling

{% hint style="info" %}

* Examples and Terraform modules should contain documentation explaining features and how to use them.
* All links in README.md files should be absolute to make Terraform Registry website show them correctly.
* Documentation may include diagrams created with [mermaid](https://github.com/mermaid-js/mermaid) and blueprints created with [cloudcraft.co](https://cloudcraft.co).
* Use [Terraform pre-commit hooks](https://github.com/antonbabenko/pre-commit-terraform) to make sure that the code is valid, properly formatted, and automatically documented before it is pushed to git and reviewed by humans.
  {% endhint %}

## Formatting

Terraform’s `terraform fmt` command enforces the canonical style for configuration files. The tool is intentionally opinionated and non-configurable, guaranteeing a uniform format across codebases so reviewers can focus on substance rather than style. Integrate it with [Terraform pre-commit hooks](https://github.com/antonbabenko/pre-commit-terraform) to validate and format code automatically before it reaches version control.

For example:

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/antonbabenko/pre-commit-terraform
    rev: v1.99.4
    hooks:
      - id: terraform_fmt
```

In CI pipelines, use `terraform fmt -check` to verify compliance. It exits with status 0 when all files are correctly formatted; otherwise, it returns a non-zero code and lists the offending files. Centralizing formatting in this way removes merge friction and enforces a consistent standard across teams.

## Editor Configuration

* **Use `.editorconfig`**: [EditorConfig](https://editorconfig.org/) helps maintain consistent coding styles for multiple developers working on the same project across various editors and IDEs. Include an `.editorconfig` file in your repositories to maintain consistent whitespace and indentation.

**Example `.editorconfig`:**

```editorconfig
[*]
indent_style = space
indent_size = 2
trim_trailing_whitespace = true

[*.{tf,tfvars}]
indent_style = space
indent_size = 2

[Makefile]
indent_style = tab
```

## Documentation

### Automatically generated documentation

[pre-commit](https://pre-commit.com/) is a framework for managing and maintaining multi-language pre-commit hooks. It is written in Python and is a powerful tool to do something automatically on a developer's machine before code is committed to a git repository. Normally, it is used to run linters and format code (see [supported hooks](https://pre-commit.com/hooks.html)).

With Terraform configurations `pre-commit` can be used to format and validate code, as well as to update documentation.

Check out the [pre-commit-terraform repository](https://github.com/antonbabenko/pre-commit-terraform/blob/master/README.md) to familiarize yourself with it, and existing repositories (eg, [terraform-aws-vpc](https://github.com/terraform-aws-modules/terraform-aws-vpc)) where this is used already.

### terraform-docs

[terraform-docs](https://github.com/segmentio/terraform-docs) is a tool that does the generation of documentation from Terraform modules in various output formats. You can run it manually (without pre-commit hooks), or use [pre-commit-terraform hooks](https://github.com/antonbabenko/pre-commit-terraform) to get the documentation updated automatically.

### Comment style

Use `#` for comments. Avoid `//` or block comments.

**Example:**

```hcl
# This is a comment explaining the resource
resource "aws_instance" "this" {
# ...
}
```

**Section Headers**: Delimit section headers in code with `# -----` or `######` for clarity.

**Example:**

```hcl
# --------------------------------------------------
# AWS EC2 Instance Configuration
# --------------------------------------------------

resource "aws_instance" "this" {
# ...
}
```

@todo: Document module versions, release, GH actions

## Resources

1. [pre-commit framework homepage](https://pre-commit.com/)
2. [Collection of git hooks for Terraform to be used with pre-commit framework](https://github.com/antonbabenko/pre-commit-terraform)
3. Blog post by [Dean Wilson](https://github.com/deanwilson): [pre-commit hooks and terraform - a safety net for your repositories](https://www.unixdaemon.net/tools/terraform-precommit-hooks/)

# FAQ

## What are the tools I should be aware of and consider using?

* [**Terragrunt**](https://terragrunt.gruntwork.io/) - Orchestration tool
* [**tflint**](https://github.com/terraform-linters/tflint) - Code linter
* [**tfenv**](https://github.com/tfutils/tfenv) - Version manager
* [**Atmos**](https://atmos.tools/) - A modern composable framework for Terraform backed by YAML
* [**asdf-hashicorp**](https://github.com/asdf-community/asdf-hashicorp) - HashiCorp plugin for the [asdf](https://github.com/asdf-vm/asdf) version manager
* [**Atlantis**](https://www.runatlantis.io/) - Pull Request automation
* [**pre-commit-terraform**](https://github.com/antonbabenko/pre-commit-terraform) - Collection of git hooks for Terraform to be used with [pre-commit framework](https://pre-commit.com/)
* [**Infracost**](https://www.infracost.io) - Cloud cost estimates for Terraform in pull requests. Works with Terragrunt, Atlantis and pre-commit-terraform too.

## What are the solutions to [dependency hell](https://en.wikipedia.org/wiki/Dependency_hell) with modules?

Versions of resource and infrastructure modules should be specified. Providers should be configured outside of modules, but only in composition. Version of providers and Terraform can be locked also.

There is no master dependency management tool, but there are some tips to make dependency specifications less problematic. For example, [Dependabot](https://dependabot.com/) can be used to automate dependency updates. Dependabot creates pull requests to keep your dependencies secure and up-to-date. Dependabot supports Terraform configurations.

# Writing Terraform configurations

## Use `locals` to specify explicit dependencies between resources

Helpful way to give a hint to Terraform that some resources should be deleted before even when there is no direct dependency in Terraform configurations.

<https://raw.githubusercontent.com/antonbabenko/terraform-best-practices/master/snippets/locals.tf>

## Terraform 0.12 - Required vs Optional arguments

1. Required argument `index_document` must be set, if `var.website` is not an empty map.
2. Optional argument `error_document` can be omitted.

{% code title="main.tf" %}

```hcl
variable "website" {
  type    = map(string)
  default = {}
}

resource "aws_s3_bucket" "this" {
  # omitted...

  dynamic "website" {
    for_each = length(keys(var.website)) == 0 ? [] : [var.website]

    content {
      index_document = website.value.index_document
      error_document = lookup(website.value, "error_document", null)
    }
  }
}
```

{% endcode %}

{% code title="terraform.tfvars" %}

```hcl
website = {
  index_document = "index.html"
}
```

{% endcode %}

### Optional Object Attributes (Terraform 1.3+)

Use optional attributes in objects to provide default values for non-required fields:

{% code title="variables.tf" %}

```hcl
variable "database_settings" {
  description = "Database configuration with optional parameters"
  type = object({
    name               = string
    engine             = string
    instance_class     = string
    backup_retention   = optional(number, 7)
    monitoring_enabled = optional(bool, true)
    tags               = optional(map(string), {})
  })
}
```

{% endcode %}

## Managing Secrets in Terraform

Secrets are sensitive data that can be anything from passwords and encryption keys to API tokens and service certificates. They are typically used to set up authentication and authorization for cloud resources. Safeguarding these sensitive resources is crucial because exposure could lead to security breaches. It’s highly recommended to avoid storing secrets in Terraform config and state, as anyone with access to version control can access them. Instead, consider using external data sources to fetch secrets from external sources at runtime. For instance, if you’re using AWS Secrets Manager, you can use the `aws_secretsmanager_secret_version` data source to access the secret value. The following example uses write-only arguments, which are supported in Terraform 1.11+, and keep the value out of Terraform state.

{% code title="main.tf" %}

```hcl
# Fetch the secret’s metadata
data "aws_secretsmanager_secret" "db_password" {
  name = "my-database-password"
}

# Get the latest secret value
data "aws_secretsmanager_secret_version" "db_password" {
  secret_id = data.aws_secretsmanager_secret.db_password.id
}

# Use the secret without persisting it to state
resource "aws_db_instance" "example" {
  engine         = "mysql"
  instance_class = "db.t3.micro"
  name           = "exampledb"
  username       = "admin"

  # write-only: Terraform sends it to AWS then forgets it
  password_wo = data.aws_secretsmanager_secret_version.db_password.secret_string
```

{% endcode %}

## Variable Validation and Input Handling

{% hint style="info" %}
Variable validation helps catch errors early, provides clear feedback, and ensures inputs meet your requirements.
{% endhint %}

### Basic Variable Validation

Use validation blocks to ensure variables meet specific criteria:

{% code title="variables.tf" %}

```hcl
variable "environment" {
  description = "Environment name for resource tagging"
  type        = string
  default     = "dev"

  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Environment must be one of: dev, staging, prod."
  }
}
```

{% endcode %}

### Object and List Validation

Validate complex data structures to ensure they contain expected values:

{% code title="variables.tf" %}

```hcl
variable "database_config" {
  description = "Database configuration"
  type = object({
    engine            = string
    instance_class    = string
    allocated_storage = number
  })

  validation {
    condition     = contains(["mysql", "postgres"], var.database_config.engine)
    error_message = "Database engine must be either 'mysql' or 'postgres'."
  }
}

variable "allowed_cidr_blocks" {
  description = "List of CIDR blocks allowed to access resources"
  type        = list(string)

  validation {
    condition = alltrue([
      for cidr in var.allowed_cidr_blocks : can(cidrhost(cidr, 0))
    ])
    error_message = "All CIDR blocks must be valid IPv4 CIDR notation."
  }
}
```

{% endcode %}
