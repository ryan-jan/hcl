# BARRY Fork of HCL Repo

This is a fork of the hashicorp/hcl repo for use with the ryan-jan/barry
project.

## Useful commands

To fix the require statements use the following.

```bash
gsed -i 's@github.com/hashicorp/hcl/v2@github.com/ryan-jan/hcl@g' ./**/*.go 
gsed -i 's@github.com/hashicorp/hcl/v2@github.com/ryan-jan/hcl@g' ./**/go.mod
```
