# Contributing Guidelines

Thanks for your interest in contributing to this project! Here are some guidelines to follow:

# Reporting Bugs

Before creating a bug report, please check the issue tracker to see if the bug has already been reported. If it has, add any additional details in the comments.

To report a new bug, open a new issue and include as much detail as possible, including:

- Steps to reproduce the bug
- Expected behavior
- Actual behavior
- Screenshots, if applicable
- Code sample that highlights the bug as well as include necessary information such as OS version and Go version

Report that are lack of information above may be autoclosed with link to this guideline.

# Suggesting Enhancements

Enhancement suggestions are welcome! To suggest an enhancement:

- Open a new issue
- Describe your suggested enhancement and explain why you think it would be useful
- Provide examples, screenshots, or anything else that could help convey the idea

# Pull Requests

Pull requests are greatly appreciated! Non code related changes such as documentation, typo fixes or bug fixes are also welcome and often accepted immediatelly

Code related changes, including feature addition will require proper discussion with maintainers first to avoid unnecessary work.

To submit a pull request:

- Fork the repository and create a new branch
- Make your changes and test them thoroughly
- Cover your changes in unit tests
- Open a pull request with a name and description of what you changed, please fill pull request issue template
- Be sure to link to any relevant issues in your PR description
- Make sure github workflow test pass
- After PR approved squash PR into 1 commit and include issue number in commit message

# Community Bug Report and Code Review

We are encourage you to review other's pull requests. Anyone with basic Go can review bug reports and pull request. You dont need to be expert to contribute to this project.

Be Constructive: Before you begin to review, remember that you are reviewing someone else's hard work. It is very easy to misunderstand something, or to have different opinions on the same topic. Remember that the goal of code review is to improve the overall quality of the code base. Always be respectful when giving code review

Here is some criteria when reviewing pull request or bug report

### Completeness Bug Report or Pull Request

Make sure bug report or pull request have basic required information based on guideline to open bug report or pull request

### Reproduce Bug or Issue
Try to reproduce bug based on step to reproduce information given by reporter. And update bug report according to your findings

### Review The Code
Read the code of pull request and check it agaisnt some common criteria:
- Does the code address the issue the PR is intended to fix/implement?
- Does the PR are within scope to solve only for that issue?
- Does it handle edge cases properly?
- Is the code easy to understand?
- Are complex sections commented for clarity?
- Does it handle errors gracefully and recover well?
- Do tests cover different cases including edge cases?
- Is the code documented clearly for future maintainers?

### Test The Code
You may want to test code on your system by cloning PR repository and test functionality wheter it already meet criteria or not

# Development Setup

To set up a development environment to work on this project:

- Clone the repository
- Run `go mod install` to install dependecies
- Create a new branch for your changes
- Make your changes and thoroughly test them
- Submit a pull request!

Thanks again for your interest in contributing!
