import axios from 'axios';
import * as Constants from '../constants/constants';
// Next we make an 'instance' of it
const instance = axios.create({
  // .. where we make our configurations
  baseURL: Constants.apiDomain,
  headers: {"Access-Control-Allow-Origin": "*",
  "Content-Type":"application/json"},
  timeout: 2000,
});

export default instance;
