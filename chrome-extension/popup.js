'use strict';
const shortCurrTab = document.getElementById('short-curr-tab');
const shortedLink = document.getElementById('result-url');
const clearButton = document.getElementById('clear-url');
const shortUrlButton = document.getElementById('short-the-url');
const inputElement = document.getElementById('original-url');
const copyButton = document.getElementById('copy-url');

copyButton.onclick = () => {
    navigator.clipboard.writeText(shortedLink.getAttribute('data-href'))
        .then(() => {console.log('copied')});
};

shortCurrTab.onclick = () => {
    chrome.tabs.query(
        {
            'active': true,
            'lastFocusedWindow': true
        },
        (tabs) => {
            inputElement.value = tabs[0].url;
            processShortLink(inputElement);
        }
    );
};

clearButton.onclick = () => {
    inputElement.value = null;
    shortedLink.setAttribute('data-href', null);
    shortedLink.innerText = null;
};

shortedLink.onclick = (element) => {
    chrome.tabs.create({'url': element.target.getAttribute('data-href'), 'active': true})
};

shortUrlButton.onclick = () => {
    if (inputElement && !inputElement.value) {
        return;
    }

    processShortLink(inputElement);
};

function processShortLink(urlElem) {
    chrome.storage.sync.get(['routes'], (result) => {
        const req = new Request(result.routes.addNewShortUrl, {
            method: 'POST',
            body: JSON.stringify({'url': urlElem.value})
        });

        getShortLink(req).then((linkResponse) => {
            if (linkResponse && !linkResponse.hasOwnProperty('shortLink')) {
                return;
            }

            let url = result.routes.goToShortUrl + linkResponse.shortLink;
            shortedLink.setAttribute('data-href', url);
            shortedLink.innerText = url;
        })
    });
}

async function getShortLink(req) {
    let response = await window.fetch(req);

    return await response.json();
}