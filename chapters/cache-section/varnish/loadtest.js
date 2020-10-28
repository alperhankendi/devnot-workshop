//k6 run loadtest.js -d 10s --vus 200

import http from 'k6/http';
import { Counter } from 'k6/metrics';
import { sleep,check, randomSeed } from 'k6';

export const requests = new Counter('http_reqs');

export const options = {
    stages: [
        { target: 20, duration: '1m' },
        { target: 15, duration: '1m' },
        { target: 0, duration: '1m' },
    ],
    thresholds: {
        requests: ['count < 100'],
    },
};

export default function () {

    for (var id = 300; id <= 400; id++) {
       let url = `http://127.0.0.1:8080/${id}`
       let res= http.get(url);

        const checkRes = check(res, {
            'status is 200': (r) => r.status === 200,
        });
    }

    sleep(1);

}