chrome.runtime.onInstalled.addListener(() => {
    let backend = {
        'routes': {
            'goToShortUrl': 'http://127.0.0.1:8000/',
            'addNewShortUrl': 'http://127.0.0.1:8000/links'
        }
    };
    chrome.storage.sync.set(backend);

    chrome.declarativeContent.onPageChanged.removeRules(undefined, function() {
        chrome.declarativeContent.onPageChanged.addRules([{
            conditions: [new chrome.declarativeContent.PageStateMatcher({
                pageUrl: {schemes: ['http', 'https']},
            })
            ],
            actions: [new chrome.declarativeContent.ShowPageAction()]
        }]);
    });
});