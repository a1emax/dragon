# Dragon

**[Youngine](https://github.com/a1emax/youngine) demo project**

## Layout

Order of packages reflects possible dependencies between them - lower packages may depend on upper ones,
but not vice versa.

* **res** - embeddable file system containing resources (assets, configs, etc).
* **pkg** - imported packages.
  * **domain** - domain logic.
  * **global** - global entities.
    * **vars** - arbitrary variables.
    * **tools** - Youngine tools, logger, randomizer, etc.
    * **assets** - static assets.
  * **window** - GUI.
  * **kernel** - control kernel.
* **cmd** - compilable service packages (if any).
* **app** - compilable application packages.
  * **desktop** - main for Windows, Linux and macOS.
  * **android_intern** - library for Android (compiled to AAR).
  * **android** - Android Studio project.

## Concept

* Game is real time.
* Level map is presented in top view and consists of cells (tiles).
* Landscape of each cell is floor or magma.
* Cell may contain wall, small wall, door, or treasure.
* Player-controlled dragon and computer-controlled thieves move around map.
* Dragon consists of segments occupying one cell each - head (one segment), body and tail (several segments).
* Thief occupies one cell.
* Walls, stones, and doors are impassable for dragon, as well as himself (he cannot move backwards and can block himself).
* Walls, magna and dragon's body are impassable for thief.
* Thief enters through door, goes to treasure, takes it, goes to door and exits through it with treasure.
* Dragon can eat the thief - to do this, his head and thief must be on the same cell. Doing this, dragon becomes longer by one segment.
* If eaten thief was carrying treasure, it drops on cell where thief was eaten. If there is already treasure in cell, dropped treasure is lost.
* Thief can chop dragon's tail — to do this, he and tail segment must be on the same cell. When it happens, segments of dragon's tail from damaged to the last disappear.
* ~~Number of thieves' enters is limited. If it is exhausted, but treasures remains, dragon wins.~~
* If thieves took all treasures through doors, dragon lost.
* If dragon has blocked himself, meaning it cannot move, he lost.
* ~~One star is given for winning itself.~~
* ~~Two stars are given if dragon protected all treasures (that is, ate all thieves and did not lose anything).~~
* ~~Three stars are given if dragon has maximum available length at game's end (that is, thieves have never chopped his tail).~~
