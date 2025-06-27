import exec from 'k6/execution';
import http from 'k6/http';
import { config } from "./config.js";

const userIds = [
  'a1927cb1-1db0-4b18-91ed-578559ba7489',
  'a7dacaa5-baa2-4458-b62e-27a5a073dfb1',
  'bb8a677a-c9f3-46ca-8299-54b23d2c4d23',
]

const promotionId = '32844198-29cc-47b6-ae53-b86f8c81be67'

export const options = {
  scenarios: {
    collect_promotion: {
      executor: 'per-vu-iterations',
      vus: userIds.length,
      iterations: 1,
      exec: 'collectPromotion'
    }
  }
}

export function collectPromotion() {
  const collectPromotionUrl = config.promotionUrl + `/api/v1/promotions/collect/${promotionId}`
  const userId = userIds[exec.vu.idInTest - 1];
  const params = {
    headers: {
      'Content-Type': 'application/json',
      'X-User-Id': userId,
    },
  }

  const resp = http.post(collectPromotionUrl, null, params);
  console.log(`Promotion ${promotionId} collected for user ${userId}: ${resp.status} - ${resp.body}`);
}