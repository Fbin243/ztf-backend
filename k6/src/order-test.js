import exec from "k6/execution";
import http from "k6/http";
import { config } from "./config.js";
import { logger } from "./utils/logger.js";

const userIds = [
  "a1927cb1-1db0-4b18-91ed-578559ba7489",
  "a7dacaa5-baa2-4458-b62e-27a5a073dfb1"
];

const merchantId = "53c7e139-5c92-49e7-a4b2-667782e8fd9e";

// Apply for all
const promotionId = "b7c9f6da-ac69-401a-a33b-379f1f565abf";

export const options = {
  scenarios: {
    pay_order_and_apply_promotion: {
      executor: "per-vu-iterations",
      vus: userIds.length,
      iterations: 1,
      exec: "payOrderAndApplyPromotion",
    },
  },
};

export function payOrderAndApplyPromotion() {
  const idx = exec.vu.idInTest - 1;
  const userId = userIds[idx];
  const params = {
    headers: {
      "Content-Type": "application/json",
      "X-User-Id": userId,
    },
  };

  // 1. Create a new order
  const createOrderUrl = config.orderUrl + "/api/v1/orders";
  const createOrderInput = JSON.stringify({
    merchant_id: merchantId,
    amount: 45000,
    info: "Cafe sữa Highland size S"
  });

  const createResp = http.post(createOrderUrl, createOrderInput, params);
  logger.info(
    `Order created for user ${userId}: ${createResp.status} - ${createResp.body}`
  );
  const orderId = JSON.parse(createResp.body).id;

  // 2. Pay for the order and apply promotion
  const payOrderUrl =
    config.orderUrl + `/api/v1/orders/pay/${orderId}`;

  const payOrderInput = JSON.stringify({
    amount: 45000,
    promotion_amount: 10000,
    pay_amount: 35000,
    promotion_id: promotionId
  });

  const payResp = http.put(payOrderUrl, payOrderInput, params);
  logger.info(
    `Promotion ${promotionId} collected for user ${userId}: ${payResp.status} - ${payResp.body}`
  );
}
