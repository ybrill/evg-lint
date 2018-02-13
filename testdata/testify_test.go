package testdata

import "github.com/stretchr/testify/suite"

type verySuite struct {
	suite.Suite
}

func (s *verySuite) TeardownSuite() { // MATCH /Testify method was spelled 'TeardownSuite', did you mean 'TearDownSuite'?/

}

// Testify doesn't care whether or not the receiver is a pointer
func (verySuite) TeardownTest() { // MATCH /Testify method was spelled 'TeardownTest', did you mean 'TearDownTest'?/

}

func (verySuite) SetUpTest() { // MATCH /Testify method was spelled 'SetUpTest', did you mean 'SetupTest'?/

}

var doesntCrashOnAnonFunctions = func() { // ok

}
