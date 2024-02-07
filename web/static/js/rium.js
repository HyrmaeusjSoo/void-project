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

const formatDateTime = date => {
    const y = date.getFullYear().toString().padStart(4, '0');
    const m = (date.getMonth() + 1).toString().padStart(2, '0');
    const d = date.getDate().toString().padStart(2, '0');
    const h = date.getHours().toString().padStart(2, '0');
    const i = date.getMinutes().toString().padStart(2, '0');
    const s = date.getSeconds().toString().padStart(2, '0');
    return `${y}-${m}-${d} ${h}:${i}:${s}`;
}

const renderJson = j => {
    typeof j === 'string' || (j = JSON.stringify(j, undefined, 4));
    j = j.replace(/&/g, '&').replace(/</g, '<').replace(/>/g, '>');
    let h = j.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, m => {
        let t = 'number';
        if (/^"/.test(m)) {
            if (/:$/.test(m)) {
                t = 'key';
            } else {
                t = 'string';
            }
        } else if (/true|false/.test(m)) {
            t = 'boolean';
        } else if (/null/.test(m)) {
            t = 'null';
        }
        return '<span class="' + t + '">' + m + '</span>';
    });
    let pre = document.createElement('pre');
    pre.className = 'jsonData';
    pre.innerHTML = h;
    return pre;
}

const object2UrlQuery = data => {
    let r = [];
    Object.entries(data).forEach(([k, v]) => {
        v.constructor == Array ? v.forEach(sv => r.push(k + "=" + sv)) : r.push(k + '=' + v);
    });
    return `?${r.join('&')}`;
}

const getCookie = n => {
    let matches = document.cookie.match(new RegExp("(?:^|; )" + n.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"));
    return matches? decodeURIComponent(matches[1]) : undefined;
}

const face = _ => {
    $c = document.createElement("link");
    $c.rel = "icon";
    $c.href = "https://HyrmaeusjSoo.github.io/CHAOS_Project/1001.ico";
    document.head.append($c);
    document.title = "javascript:;";
}

const webkitFilter = (f = 'grayscale(100%)') => {
    document.body.style.webkitFilter = f;
}
