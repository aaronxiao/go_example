package dataStruct

import (
	"log"
	"testing"
)

func TestLeftRotate(t *testing.T){
	var i10 Int = 10
	var i12 Int = 12

	rbTree := New()

	x := &Node{rbTree.NIL, rbTree.NIL, rbTree.NIL, BLACK, i10}
	rbTree.root = x
	y := &Node{rbTree.root.Right, rbTree.NIL, rbTree.NIL, RED, i12}
	rbTree.root.Right = y

	log.Println("root : ", rbTree.root)
	log.Println("left : ", rbTree.root.Left)
	log.Println("right : ", rbTree.root.Right)

	rbTree.LeftRotate(rbTree.root)

	log.Println("root : ", rbTree.root)
	log.Println("left : ", rbTree.root.Left)
	log.Println("right : ", rbTree.root.Right)

}


func TestInsert(t *testing.T){
	rbTree := New()

	rbTree.Insert(&Node{rbTree.NIL, rbTree.NIL, rbTree.NIL, RED, Int(10)})
	rbTree.Insert(&Node{rbTree.NIL, rbTree.NIL, rbTree.NIL, RED, Int(9)})
	rbTree.Insert(&Node{rbTree.NIL, rbTree.NIL, rbTree.NIL, RED, Int(8)})
	rbTree.Insert(&Node{rbTree.NIL, rbTree.NIL, rbTree.NIL, RED, Int(6)})
	rbTree.Insert(&Node{rbTree.NIL, rbTree.NIL, rbTree.NIL, RED, Int(7)})

	log.Println("rbtree counts : ", rbTree.count)

	log.Println("------ ", rbTree.root.Item)
	log.Println("----", rbTree.root.Left.Item, "---", rbTree.root.Right.Item)
	log.Println("--", rbTree.root.Left.Left.Item, "-", rbTree.root.Left.Right.Item)

}


func TestRightRotate(t *testing.T){
	var i10 Int = 10
	var i12 Int = 12

	rbtree := New()

	x := &Node{rbtree.NIL, rbtree.NIL, rbtree.NIL, BLACK, i10}
	rbtree.root = x
	y := &Node{rbtree.root.Right, rbtree.NIL, rbtree.NIL, RED, i12}
	rbtree.root.Left = y

	log.Println("root : ", rbtree.root)
	log.Println("left : ", rbtree.root.Left)
	log.Println("right : ", rbtree.root.Right)

	rbtree.RightRotate(rbtree.root)

	log.Println("root : ", rbtree.root)
	log.Println("left : ", rbtree.root.Left)
	log.Println("right : ", rbtree.root.Right)

}

func TestItem(t *testing.T){
	var iType1 Int = 10
	var iType2 Int = 12

	log.Println(iType1.Less(iType2))


	var strType1 String = "sola"
	var strType2 String = "ailumiyana"

	log.Println(strType1.Less(strType2))
}
