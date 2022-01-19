
// Comparators are used to compare two objects. In this challenge, you'll create a comparator and use it to
// sort an array. The Player class is provided in the editor below. It has two fields:

// 1.name: a string.
// 2.score: an integer.

class Checker implements Comparator<Player> {
    // complete this method
  public int compare(Player a, Player b) {
      if (a.score < b.score) {
          return 1;
      } else if (a.score == b.score) {            
          int order = a.name.compareTo(b.name);
          return order;
      } else {
          return -1;
      }
  }
}
