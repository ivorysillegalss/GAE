import axios from 'axios'
import * as URL from './url'

export function getUserInfo(params: any) {
    return axios.get(URL.USER_INFO,{params})
}

export function getUserSpecificInfo(params: any) {
    return axios.get(URL.USER_SPECIFIC_INFO,{params})
}

export function getRankLevel(params: any) {
    return axios.get(URL.RANK_LEVEL,{params})
}

export function getRankTech(params: any) {
    return axios.get(URL.RANK_TECH,{params})
}

export function getRankNation(params: any) {
    return axios.get(URL.RANK_NATION,{params})
}

export function getRankHotPhase() {
    return axios.get(URL.RANK_HOT_PHASE)
}

export function getRankHot(params: any) {
    return axios.get(URL.RANK_HOT,{params})
}

export function getRankInit() {
    return axios.get(URL.RANK_INIT)
}

export function getRelation(params: any) {
    return axios.get(URL.RELATION,{params})
}