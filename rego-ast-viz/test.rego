default matches_sgx_policy = false
matches_sgx_policy = true {
  input.sgx_mrenclave == "1234e819861adef61232a4865efea9337b91ed301233491b17f123d9e8212311"
  input.sgx_mrsigner == "123719e77d123a1470f612362a4d774303c899db69123f9c70ee123c08123123"
  input.sgx_is_debuggable == false
}