import { config } from "./config.js";
// Import the http module to make HTTP requests. From this point, you can use `http` methods to make HTTP requests.
import http from 'k6/http';

// Import the sleep function to introduce delays. From this point, you can use the `sleep` function to introduce delays in your test script.
import { sleep } from 'k6';

export const options = {
  iterations: 10,
}

export default function () {
  // Make a GET request to the target URL
  http.get(config.orderUrl + "/health");

  // Sleep for 1 second to simulate real-world usage
  sleep(1);
}