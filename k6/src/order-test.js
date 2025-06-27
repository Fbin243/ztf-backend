import { config } from "./config.js";
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  iterations: 10,
}

export default function () {
  const payload = JSON.stringify({

  })
  http.get(config.orderUrl + "/health");

  sleep(1);
}