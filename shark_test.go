package hunt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	shark := Shark{hungry:true,tired:false,speed: 20}
	prey := Prey{speed: 10}

	err:=shark.Hunt(&prey)
	
	assert.NoError(t, err)
	assert.Nil(t, err)

	err = shark.Hunt(&prey)
	
	assert.Error(t, err)
	assert.Equal(t, true,shark.tired)
	assert.Equal(t, false, shark.hungry)
	
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//arrange
	shark := Shark{tired: true}
	prey := Prey{name: "prey"}
	IsTired := errors.New("cannot hunt, i am really tired")
	//act 
	err := shark.Hunt(&prey)
	assert.Error(t, err)
	assert.Equal(t,err,IsTired)
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	shark := Shark{hungry: false}
	prey := Prey{name: "prey"}
	IsHungry := errors.New("cannot hunt, i am not hungry")

	err := shark.Hunt(&prey)
	assert.Error(t, err)
	assert.Equal(t, IsHungry, err)

}

func TestSharkCannotReachThePrey(t *testing.T) {
	shark := Shark{hungry:true,tired:false,speed: 10}
	prey := Prey{speed: 20}
	isFast := errors.New("could not catch it")

	err:=shark.Hunt(&prey)
	assert.Error(t, err)
	assert.Equal(t, isFast, err)
}

func TestSharkHuntNilPrey(t *testing.T) {
}
