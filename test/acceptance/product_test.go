package acceptance_test

func (suite *AcceptanceSuite) TestCreateProduct() {
	createdProduct, err := suite.CreateProduct()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(createdProduct)

	suite.AssertProductAppearsInSearchResults(createdProduct)
}
