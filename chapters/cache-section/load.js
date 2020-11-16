import http from 'k6/http';
import { check } from 'k6';

export default function () {

    let batch_urls = []
    for (let x=100;x<200; x++)
        batch_urls.push(['GET','http://127.0.0.1:8080/'+x])

    let responses = http.batch(batch_urls)
    check(responses[0], {
        'main page status was 200': (res) => res.status === 200,
    });
}

