# To Do

[ ] figure out how to deal with events
  [ ] need to build hierarchy from the outside (certain components have children but don't expose them)
  [ ] widgets somehow need to communicate that they are interested in certain events
[ ] make (certain?) widgets focusable
[ ] add a button type
[ ] add support for padding/margins
[ ] don't render widgets if bounds are 0

# Done

[x] move dimensions and bounds related code into reusable type
[x] add border type

# Ideas

# Thoughts

## Mouse Event Dispatching

A mouse click (or any other event that can be anywhere) happens.
Ideally we dispatch some kind of mouse click event to the component closest to that event.
Closest component would be the one that encases the event point and is at the top of the component stack.
The closest component might or might not care about this kind of event. If it doesn't, we ask the parent.

