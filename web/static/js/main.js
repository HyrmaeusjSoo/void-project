
const request = async (url, method='GET', data=null, contentType='application/json;charset=utf-8') => {
    let res;
    let option = {
        method: method,
        headers: {
            'userId': localStorage.getItem('userId'),
            'token': localStorage.getItem('token'),
        },
    }
    method != 'GET' && (option.headers['Content-Type'] = contentType, option.body = JSON.stringify(data));
    await fetch(url, option).then(r => r.json()).then(r => res = r);
    return res;
}
