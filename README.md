# Design Patterns with Golang

Reference: <https://refactoring.guru/design-patterns>

Here some quick notes

## Behavorial Patterns

### Observer

- Things that notify others about their changes in state should implement the publisher interface
- Things that want to track changes to the publisher's state should implement the subscriber interface

The publisher interface includes:
- an array storing references to subscriber objects
- method to add/remove subscriber from above list

The subscriber interface includes:
- the notification method

Others
- all subscribers should implement the same interface
- and the publisher communicate with them via that only interface
- if you have many subscriber interface and want subcribers to be compatible with them, you can go further and make those subsriber implement one interface
