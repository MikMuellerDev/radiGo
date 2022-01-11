function startSpinner(spinnerId) {
    const spinner = document.getElementById(spinnerId)
    if (spinner.style.opacity !== '100%')
        if (spinner.style.animation !== 'spinnerAnimation 1s linear infinite')
            spinner.style.animation = 'spinnerAnimation 1s linear infinite'
    spinner.style.opacity = '100%'
}

function stopSpinner(spinnerId) {
    const spinner = document.getElementById(spinnerId)
    if (spinner.style.opacity === '1') {
        setTimeout(function () {
            spinner.style.opacity = '0'
            setTimeout(function () {
                spinner.style.animation = 'none'
                spinner.style.opacity = 'none'
            }, 200)
        }, 800)
    }
}