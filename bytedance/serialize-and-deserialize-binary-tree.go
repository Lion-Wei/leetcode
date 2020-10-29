package bytedance

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{
		l: make([]string, 0),
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return strings.Join(mySer(root), ",")
}

func mySer(n *TreeNode) []string {
	if n == nil {
		return []string{"nil"}
	}
	res := []string{fmt.Sprintf("%d", n.Val)}
	res = append(res, mySer(n.Left)...)
	res = append(res, mySer(n.Right)...)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.l = strings.Split(data, ",")
	return this.myDe()
}

func (this *Codec) myDe() *TreeNode {
	if this.l[0] == "nil" {
		this.l = this.l[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.l[0])
	this.l = this.l[1:]
	return &TreeNode{
		Val:   val,
		Left:  this.myDe(),
		Right: this.myDe(),
	}
}
