import axios from 'axios'

/**
 * Tiny, in-line analytics helper.
 * – No worker, no queue.
 * – Never throws (errors are swallowed and logged).
 */

const ENDPOINT = '/analytics/events'   // backend

/** Detects device class from the UA string. */
function getDevice () {
  const ua = navigator.userAgent
  if (/Mobi|Android/i.test(ua)) return 'mobile'
  if (/Tablet|iPad/i.test(ua)) return 'tablet'
  return 'desktop'
}

/** Low-level send — single POST per event. */
function send (event, data) {
  return axios.post(ENDPOINT, { event, ...data }).catch(err => {
    console.warn('[analytics] failed', err)
  })
}

/** Public helper: call `track('event_name', { ...payload })` anywhere. */
export function track (event, payload = {}) {
  return send(event, { device: getDevice(), ...payload })
}

/** One-call installer used in main.js for automatic page tracking. */
export function installAnalytics (router) {
  router.afterEach((to, from) => {
    track('page_visited', {
      path: to.fullPath,
      referrer: from.fullPath || document.referrer || null
    })
  })
}
