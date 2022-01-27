(function () {
    // Fake html meta tag to disable darkreader.
    const fakeMetaTag = document.createElement("meta")
    fakeMetaTag.name = "darkreader"
    fakeMetaTag.content = "off"

    // Alter the real metatag with the fake one.
    const alterMetaTag = function () {
        let correctTag = document.querySelector(
            'meta[content="' + fakeMetaTag.content + '"]'
        )
        if (!correctTag) {
            document.head.appendChild(fakeMetaTag)
        }
        let realTag: HTMLMetaElement | null = document.querySelector(
            'meta[name="' + fakeMetaTag.name + '"]'
        )
        if (realTag && realTag.content != fakeMetaTag.content) {
            realTag.remove()
        }
    }

    // Remove all Darkreader style tags form `document.head`.
    const removeDarkreader = function () {
        // NOTE: use traditional 'for loops' for IE 11
        for (const style of document.head.getElementsByClassName("darkreader")) {
            style.remove()
        }
    }

    // Observing callback function.
    const callback = function () {
        alterMetaTag()
        removeDarkreader()
    }

    // Options for the observer (which mutations to observe).
    const config = { attributes: false, childList: true, subtree: false }

    // Create an observer instance linked to the callback function.
    const observer = new MutationObserver(callback)

    if (
        !document.querySelector('meta[content="' + fakeMetaTag.content + '"]') &&
        document.querySelector('meta[name="' + fakeMetaTag.name + '"]')
    ) {
        console.error(
            "Please add the line bellow to your index.html:\n",
            '<meta name="darkreader" content="off">\n',
            "or you may encounter performance issues!\n",
            "\nplease take a look at: https://github.com/hadialqattan/no-darkreader#usage"
        )
    } else {
        // Start observing the target node for configured mutations.
        observer.observe(document.head, config)

        // Execute for the fist time to take effect.
        callback()
    }
})()
