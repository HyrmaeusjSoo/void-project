/*
┌───────────────────────────────────────────────────────────────────────────────────────────┐
│ Sū Shēngxǜ's from past to present VOID CHAOS False Philosophy code.
├───────────────────────────────────────────────────────────────────────────────────────────┤
│ Elysium, in the Blue Sky. ファンタジーアドベンチャー。 泡泡枪、七彩、环世界宇宙飞船
├───────────────────────────────────────────────────────────────────────────────────────────┤
│ 银河系 🌌⚛️🔮🗡️✡️🏞️🎮 Requests.                                              2023-2024
├───────────────────────────────────────────────────────────────────────────────────────────┤
│                                                                    —————— Hyrmaeusj 苏
└───────────────────────────────────────────────────────────────────────────────────────────┘
*/

const host = `http://${window.location.host}/api/v1/`;
const request = async (url, method='GET', data=null, contentType='application/json;charset=utf-8') => {
    let res;
    let option = {
        method: method,
        headers: {
            'user_id': localStorage.getItem('user_id'),
            'token': localStorage.getItem('token'),
        },
    }
    method != 'GET' && (option.headers['Content-Type'] = contentType, option.body = JSON.stringify(data));
    await fetch(host+url, option).then(r => r.json()).then(r => res = r);
    return res;
};
const postForm = async (url, formData) => {
    let res;
    let option = {
        method: 'POST',
        headers: {
            'user_id': localStorage.getItem('user_id'),
            'token': localStorage.getItem('token'),
        },
    }
    option.body = formData
    await fetch(host+url, option).then(r => r.json()).then(r => res = r);
    return res;
};
