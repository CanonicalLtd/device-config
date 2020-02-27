/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

const API_PREFIX = '/v1/'

function getBaseURL() {
    return window.location.protocol + '//' + window.location.hostname + ':' + window.location.port + API_PREFIX;
}

let Constants = {
    baseUrl: getBaseURL(),
    LoadingImage: '/static/images/ajax-loader.gif',
    //missingIcon: '/static/images/snapcraft-missing-icon.svg',
}

export default Constants