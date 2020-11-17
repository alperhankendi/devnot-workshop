import http from 'k6/http';
import {  check } from 'k6';

function randomString(length) {
    const charset = 'abcdefghijklmnopqrstuvwxyz0123456789';
    let res = '';
    while (length--) res += charset[Math.random() * charset.length | 0];
    return res;
}

export default function () {

    var url = 'http://127.0.0.1:5001/api/v1/v3';

    var pl = {
        "title": "Miss Jerry",
        "original_title": "Miss Jerry",
        "year": 1894,    "date_published": "1894-10-09",    "genre": "Romance",
        "duration": 45,    "country": "USA",
        "director": "Alexander Black",    "writer": "Alexander Black",
        "production_company": "Alexander Black Photoplays",
        "actors": "Blanche Bayliss, William Courtenay, Chauncey Depew",
        "description": "The adventures of a female reporter in the 1890s.",
        "votes": 154}
        pl["imdb_title_id"]= randomString(20)
    var payload = JSON.stringify(
            pl
    );
    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };
    const res = http.post(url, payload, params);
    const checkRes = check(res, {
        'status is 200': (r) => r.status === 201,
    });
}
