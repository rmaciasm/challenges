func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	return (yourLeft == friendsLeft && yourRight == friendsRight) ||
		(yourLeft == friendsRight && yourRight == friendsLeft)
}
   