package myfunction

import (
	"strings"
	"testing"
)

const someText = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, " +
	"who are you not to be? Your playing small does not serve the world. There is nothing enlightened " +
	"about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. " +
	"We were born to make manifest the glory that is within us. It's not just in some of us; " +
	"it's in everyone. And as we let our own light shine, " +
	"we unconsciously give other people permission to do the same. " +
	"As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

//testing

func TestJointextManualWay(t *testing.T) {
	d := strings.Split(someText, " ")
	s := JointextManualWay(d)
	if s != someText {
		t.Errorf("error! unexpected : %s", s)
	}
}

func TestJointextFastWay(t *testing.T) {
	d := strings.Split(someText, " ")
	s := strings.Join(d, " ")
	if s != someText {
		t.Errorf("error! unexpected : %s", s)
	}
}

//benchmarking

var z []string

func BenchmarkJointextManualWay(b *testing.B) {
	z := strings.Split(someText, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JointextManualWay(z)
	}
}

func BenchmarkJoinTextFastWay(b *testing.B) {
	z := strings.Split(someText, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JoinTextFastWay(z)
	}
}
