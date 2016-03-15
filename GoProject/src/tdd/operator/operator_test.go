package operator
import(
  "github.com/stretchr/testify/assert"
  "testing"
)
const(
  num1 = 1
  num2 = 2
)
func TestPlus(t *testing.T){
  assert.Equal(t, Plus(num1,num2), 3)
}
func TestMinus(t *testing.T){
  assert.Equal(t,Minus(num1,num2),-1)
}
func TestMultipy(t *testing.T){
  assert.Equal(t,Multipy(num1,num2),2)
}
func TestDevide(t *testing.T){
  assert.Equal(t,Devide(num1,num2),0)
}
