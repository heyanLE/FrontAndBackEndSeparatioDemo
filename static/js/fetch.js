function FetchGet(url,callback) {
    fetch(url,{method : 'GET'})
        .then((result) => {
            return result.json()
        }).then((resultJson) => {
            callback(resultJson.code,resultJson.message,resultJson.value)
    })
}

function FetchPost(url,body,callback) {
    fetch(url,{
        method : 'POST',
        body : JSON.stringify(body)
    })
        .then((result) => {
            return result.json()
        }).then((resultJson) => {
        callback(resultJson.code,resultJson.message,resultJson.value)
    })
}