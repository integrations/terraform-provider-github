# Import by ID
terraform import github_repository_autolink_reference.auto my-repo/123
# See the GitHub documentation for how to list all autolinks of a repository to learn the autolink ids to use with the import command.
# https://docs.github.com/en/rest/repos/autolinks#list-all-autolinks-of-a-repository

# Import by key prefix
terraform import github_repository_autolink_reference.auto oof/OOF-
