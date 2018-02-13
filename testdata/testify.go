package testdata

import "github.com/stretchr/testify/suite"

// This is not a test, so its errors will not be reported

type someSuite struct {
	suite.Suite
}

func (*someSuite) SetUpSuite() { // OK
}
