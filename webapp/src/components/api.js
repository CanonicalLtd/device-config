/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

import axios from 'axios'
import constants from './constants'

let service = {
    loginRequest:  (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'login', query);
    },

    networkGet: () => {
        return axios.get(constants.baseUrl + 'network');
    },

    networkUpdate: (iface) => {
        return axios.post(constants.baseUrl + 'network', iface);
    },

    networkApply: () => {
        return axios.post(constants.baseUrl + 'network/apply');
    },

    proxyGet: () => {
        return axios.get(constants.baseUrl + 'proxy');
    },

    proxyUpdate: (iface) => {
        return axios.post(constants.baseUrl + 'proxy', iface);
    },

    timeGet: () => {
        return axios.get(constants.baseUrl + 'time');
    },

    timeUpdate: (t) => {
        return axios.post(constants.baseUrl + 'time', {ntp: t.ntp, time: t.time, timezone: t.timezone});
    },

}

export default service