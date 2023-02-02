package hunt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	
	
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//arrange
	shark := Shark{tired: true}
	prey := Prey{}
	IsTired := errors.New("cannot hunt, i am really tired")
	//act 
	err := shark.Hunt(&prey)
	assert.Error(t, err)
	assert.Equal(t,err,IsTired)
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
}

func TestSharkCannotReachThePrey(t *testing.T) {
}

func TestSharkHuntNilPrey(t *testing.T) {
}
