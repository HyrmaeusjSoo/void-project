
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
    $c.href = "https://hyleasoo.github.io/CHAOS_Project/1001.ico";
    document.head.append($c);
    document.title = "javascript:;";
}

const webkitFilter = (f = 'grayscale(100%)') => {
    document.body.style.webkitFilter = f;
}

/* Now between time periods */
function _isTimeBetween(beginTime, endTime) {
    let strb = beginTime.split(":");
    if (strb.length < 2) {
        return false;
    }
    let stre = endTime.split(":");
    if (stre.length < 2) {
        return false;
    }
    let b = new Date();
    let e = new Date();
    let n = new Date();
    b.setHours(strb[0]);
    b.setMinutes(strb[1]);
    e.setHours(stre[0]);
    e.setMinutes(stre[1]);
    if (n.getTime() - b.getTime() > 0 && n.getTime() - e.getTime() < 0) {
        return true;
    } else {
        return false;
    }
}

/* Replace two float */
const _replace2Float = obj => {
    obj.value = obj.value.replace(/[^\d.]/g, ""); //clear symbol
    obj.value = obj.value.replace(/^\./g, ""); //first number
    obj.value = obj.value.replace(/\.{2,}/g, "."); //first point
    obj.value = obj.value.replace(".", "$#$").replace(/\./g, "").replace("$#$", ".");
    obj.value = obj.value.replace(/^(\-)*(\d+)\.(\d\d).*$/, '$1$2.$3'); //decimal
}

/* Verify letter and number */
const checkNormalString = value => {
    let reg = /^[0-9a-zA-Z]*$/g;
    return reg.test(value);
}

/* Someone,or else */
const orelse = _ => {
    let pos = prompt('number:', '1');
    let spe1 = false || false || true;
    let spe2 = true && true && false;
    let sp1 = (pos == '1' && 'a') || (pos == '2' && 'b') || (pos == '3' && 'c') || 'n';
    let sp2 = pos == '1' && 'a' || pos == '2' && 'b' || pos == '3' && 'c' || 'n';
    let sp3 = { '1': 'a', '2': 'b', '3': 'c' }[pos] || 'n';
    return sp3;
}

/* Encode unicode */
const encodeUnicode = str => {
    let res = [];
    for (let i = 0; i < str.length; i++) {
        res[i] = ("00" + str.charCodeAt(i).toString(16)).slice(-4);
    }
    return "\\u" + res.join("\\u");
}

/* Decode unicode */
const decodeUnicode = str => unescape(str.replace(/\\/g, "%"));

/* seconds */
function seconds() {

    const arrayConcat = (arr, ...args) => arr.concat(...args);
    // arrayConcat([1], 2, [3], [[4]]) // [1,2,3,[4]]

    const difference = (a, b) => { const s = new Set(b); return a.filter(x => !s.has(x)); };
    // difference([1,2,3], [1,2]) // [3]

    const intersection = (a, b) => { const s = new Set(b); return a.filter(x => s.has(x)); };
    // intersection([1,2,3], [4,3,2]) // [2,3]

    const union = (a, b) => Array.from(new Set([...a, ...b]));
    // union([1,2,3], [4,3,2]) // [1,2,3,4]

    const average = arr => arr.reduce((acc, val) => acc + val, 0) / arr.length;
    // average([1,2,3]) // 2

    const chunk = (arr, size) => Array.from({ length: Math.ceil(arr.length / size) }, (v, i) => arr.slice(i * size, i * size + size));
    // chunk([1,2,3,4,5], 2) // [[1,2],[3,4],5]

    const compact = (arr) => arr.filter(v => v);
    // compact([0, 1, false, 2, '', 3, 'a', 'e'*23, NaN, 's', 34]) // [ 1, 2, 3, 'a', 's', 34 ]

    const countOccurrences = (arr, value) => arr.reduce((a, v) => v === value ? a + 1 : a + 0, 0);
    // countOccurrences([1,1,2,1,2,3], 1) // 3

    const deepFlatten = arr => arr.reduce((a, v) => a.concat(Array.isArray(v) ? deepFlatten(v) : v), []);
    // deepFlatten([1,[2],[[3],4],5]) // [1,2,3,4,5]

    const dropElements = (arr, func) => {
        while (arr.length > 0 && !func(arr[0])) arr.shift();
        return arr;
    };
    // dropElements([1, 2, 3, 4], n => n >= 3) // [3,4]

    const fillArray = (arr, value, start = 0, end = arr.length) => arr.map((v, i) => i >= start && i < end ? value : v);
    // fillArray([1,2,3,4],'8',1,3) // [1,'8','8',4]

    const filterNonUnique = arr => arr.filter(i => arr.indexOf(i) === arr.lastIndexOf(i));
    // filterNonUnique([1,2,2,3,4,4,5]) // [1,3,5]

    const elementIsVisibleInViewport = (el, partiallyVisible = false) => {
        const { top, left, bottom, right } = el.getBoundingClientRect();
        return partiallyVisible
            ? ((top > 0 && top < innerHeight) || (bottom > 0 && bottom < innerHeight)) &&
            ((left > 0 && left < innerWidth) || (right > 0 && right < innerWidth))
            : top >= 0 && left >= 0 && bottom <= innerHeight && right <= innerWidth;
    };
    // e.g. 100x100 viewport and a 10x10px element at position {top: -1, left: 0, bottom: 9, right: 10}
    // elementIsVisibleInViewport(el) // false (not fully visible)
    // elementIsVisibleInViewport(el, true) // true (partially visible)

    const redirect = (url, asLink = true) => asLink ? window.location.href = url : window.location.replace(url);
    // redirect('https://host.com')

    const scrollToTop = _ => {
        const c = document.documentElement.scrollTop || document.body.scrollTop;
        if (c > 0) {
            window.requestAnimationFrame(scrollToTop);
            window.scrollTo(0, c - c / 8);
        }
    };
    // scrollToTop()

    const getDaysDiffBetweenDates = (dateInitial, dateFinal) => (dateFinal - dateInitial) / (1000 * 3600 * 24);
    // getDaysDiffBetweenDates(new Date("2019-08-12"), new Date("2019-08-22")) // 10

    const reverseString = str => [...str].reverse().join('');
    // reverseString('foobar') // 'raboof'

    const sortCharactersInString = str => str.split('').sort((a, b) => a.localeCompare(b)).join('');
    // sortCharactersInString('cabbage') // 'aabbceg'

    const truncate = (str, num) => str.length > num ? str.slice(0, num > 3 ? num - 3 : num) + '...' : str;
    // truncate('boomerang', 7) // 'boom...'

    const randomIntegerInRange = (min, max) => Math.floor(Math.random() * (max - min + 1)) + min;
    // randomIntegerInRange(0, 5) // 2

    const rgbToHex = (r, g, b) => ((r << 16) + (g << 8) + b).toString(16).padStart(6, '0');
    // rgbToHex(255, 165, 1) // 'ffa501'

    const hexToRGB = hex => {
        let alpha = false,
            h = hex.slice(hex.startsWith('#') ? 1 : 0);
        if (h.length === 3) h = [...h].map(x => x + x).join('');
        else if (h.length === 8) alpha = true;
        h = parseInt(h, 16);
        return (
            'rgb' +
            (alpha ? 'a' : '') +
            '(' +
            (h >>> (alpha ? 24 : 16)) +
            ', ' +
            ((h & (alpha ? 0x00ff0000 : 0x00ff00)) >>> (alpha ? 16 : 8)) +
            ', ' +
            ((h & (alpha ? 0x0000ff00 : 0x0000ff)) >>> (alpha ? 8 : 0)) +
            (alpha ? `, ${h & 0x000000ff}` : '') +
            ')'
        );
    };
    //hexToRGB('#27ae60ff'); // 'rgba(39, 174, 96, 255)'
    //hexToRGB('27ae60'); // 'rgb(39, 174, 96)'
    //hexToRGB('#fff'); // 'rgb(255, 255, 255)'

    const getURLParameters = url => (url.match(/([^?=&]+)(=([^&]*))/g) || []).reduce(
        (a, v) => ((a[v.slice(0, v.indexOf('='))] = v.slice(v.indexOf('=') + 1)), a), {}
    );
    //getURLParameters('http://url.com/page?name=Adam&surname=Smith'); // {name: 'Adam', surname: 'Smith'}
    //getURLParameters('google.com'); // {}

    const validateEmail = str => /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(str);
    // validateEmail(mymail@gmail.com) // true

    const validateNumber = n => !isNaN(parseFloat(n)) && isFinite(n) && Number(n) == n;
    // validateNumber('10') // true

    const prettyBytes = (num, precision = 3, addSpace = true) => {
        const UNITS = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
        if (Math.abs(num) < 1) return num + (addSpace ? ' ' : '') + UNITS[0];
        const exponent = Math.min(Math.floor(Math.log10(num < 0 ? -num : num) / 3), UNITS.length - 1);
        const n = Number(((num < 0 ? -num : num) / 1000 ** exponent).toPrecision(precision));
        return (num < 0 ? '-' : '') + n + (addSpace ? ' ' : '') + UNITS[exponent];
    };
    // prettyBytes(1000); // "1 KB"
    // prettyBytes(-27145424323.5821, 5); // "-27.145 GB"
    // prettyBytes(123456789, 3, false); // "123MB"

    const isValidJSON = str => {
        try {
            JSON.parse(str);
            return true;
        } catch (e) {
            return false;
        }
    };
    //isValidJSON('{"name":"Adam","age":20}'); // true
    //isValidJSON('{"name":"Adam",age:"20"}'); // false
    //isValidJSON(null); // true

    return {
        randomIntegerInRange: randomIntegerInRange,
        validateEmail: validateEmail,
        validateNumber: validateNumber,
        prettyBytes: prettyBytes,
        isValidJSON: isValidJSON
    };

}

const _mainTme = '6307#1`2`3`';
const _Zeno = (yr, hx, xx = 5, yq) =>
    (_ =>
        zenogl
    )(e = 1 + 2 - 1
        , xfjs = (xs, fh = 2) =>
            Math.floor(xs * 10 ** fh) / 10 ** fh
        , zenogl = (l =>
            xfjs(
                ({ '\u006b\u006d\u0073': 0.1, '\u004a\u0069': 0.31, '\u004c\u0067': 1.88, '\u0042\u0063\u0071': 0.25, '\u0059\u0075\u0065': 0.13 }[yr] || 0.1)
                * ({ '\u0070\u0074': 1, '\u0078\u0079': 1.5, '\u0073\u0073': 3 }[hx] || 1)
                * xfjs(0.01 * (100 + 5 * xx))
                * (yq <= 1 && 1 || xfjs(yq))
                , 3)
        )(yq = (y => Math.sqrt(yq || 21) * 0.01 * 1.3 + 0.95)())
    )
    + ' - ' + xfjs(1 / zenogl * 100, 0)
    ;
ateEmail,
        validateNumber: validateNumber,
        prettyBytes: prettyBytes,
        isValidJSON: isValidJSON
    };

}

const _mainTme = '6307#1`2`3`';
const _Zeno = (yr, hx, xx = 5, yq) =>
    (_ =>
        zenogl
    )(e = 1 + 2 - 1
        , xfjs = (xs, fh = 2) =>
            Math.floor(xs * 10 ** fh) / 10 ** fh
        , zenogl = (l =>
            xfjs(
                ({ '\u006b\u006d\u0073': 0.1, '\u004a\u0069': 0.31, '\u004c\u0067': 1.88, '\u0042\u0063\u0071': 0.25, '\u0059\u0075\u0065': 0.13 }[yr] || 0.1)
                * ({ '\u0070\u0074': 1, '\u0078\u0079': 1.5, '\u0073\u0073': 3 }[hx] || 1)
                * xfjs(0.01 * (100 + 5 * xx))
                * (yq <= 1 && 1 || xfjs(yq))
                , 3)
        )(yq = (y => Math.sqrt(yq || 21) * 0.01 * 1.3 + 0.95)())
    )
    + ' - ' + xfjs(1 / zenogl * 100, 0)
    ;
