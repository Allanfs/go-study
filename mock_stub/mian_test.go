package mockstub_test

import (
	"testing"
	"time"

	"allanfs.com/mocks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockery_OnMethodCalled_ReturnSomething(t *testing.T) {
	ironbergMockado := mocks.NewIronberg(t)
	ironbergMockado.On("String").Return("Bolado")

	assert.Equal(t, "Bolado", ironbergMockado.String())
}

func TestMockery_OnMethodCalled_PanicSomething(t *testing.T) {
	ironbergMockado := mocks.NewIronberg(t)
	ironbergMockado.On("String").Panic("panic esperado")

	assert.Panics(t, func() { ironbergMockado.String() })
}

func TestMockery_OnMethodCalled_ReturnAfterTime(t *testing.T) {
	ironbergMockado := mocks.NewIronberg(t)

	ironbergMockado.
		On("String").
		After(time.Second * 5).
		Return("Esperoou 5 sec")

	assert.Equal(t, "Esperoou 5 sec", ironbergMockado.String())
}

func TestMockery_OnMethodCalled_ReturnNTimes(t *testing.T) {
	ironbergMockado := mocks.NewIronberg(t)

	ironbergMockado.
		On("String").
		Times(3).
		Return("Deve executar no máximo 3 vezes")

	assert.Equal(t, "Deve executar no máximo 3 vezes", ironbergMockado.String())
	assert.Equal(t, "Deve executar no máximo 3 vezes", ironbergMockado.String())
	assert.Equal(t, "Deve executar no máximo 3 vezes", ironbergMockado.String())
	// se a linha abaixo for executada, o teste falha
	// assert.Equal(t, "Deve executar no máximo 3 vezes", ironbergMockado.String())
}
