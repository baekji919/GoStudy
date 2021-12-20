package main

    import (
        "github.com/stretchr/testify/assert"
        "testing"
    )

    func TestValidateBusinessNumber_유효한경우(t *testing.T) {
        //given
        a := "2018121515"
        //when
        actual :=  ValidateBusinessNumber(a)
        //then
        assert.Equal(t, true, actual)
    }


    func TestValidateBusinessNumber_10자리아닌경우(t *testing.T) {
        //given
        a := "123456789"
        //when
        actual :=  ValidateBusinessNumber(a)
        //then
        assert.Equal(t, false, actual)
    }


    func TestValidateBusinessNumber_유효한번호아닌경우(t *testing.T) {
        //given
        a := "1234567890"
        //when
        actual :=  ValidateBusinessNumber(a)
        //then
        assert.Equal(t, false, actual)
    }
