


let start = document.getElementById("start_button");

start.onclick = function () {
    window.location.href = "/research";
}


const download_file = document.getElementById("file")

download_file.onclick = async function () {
    let url = "/download_file"
    let options = {
        method: 'GET',
        headers: new Headers({
            'Content-Type': 'application/json',
        }),
        mode: 'cors',
        cache: 'default'
    };

    //Perform a GET Request to server
    fetch(url, options)
        .then(function (response) {
            return response.blob();
        }).then(function (myBlob) {
        let downloadLink = window.document.createElement('a');
        downloadLink.href = window.URL.createObjectURL(new Blob([myBlob]));
        downloadLink.download = "research.docx";
        document.body.appendChild(downloadLink);
        downloadLink.click();
        document.body.removeChild(downloadLink);
    }).catch((e) => {
        console.log("e", e);
    });
}