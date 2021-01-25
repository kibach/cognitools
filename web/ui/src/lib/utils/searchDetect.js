const PARTIAL_EMAIL_REGEX = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@/;
const PARTIAL_PHONE_REGEX = /^\+\d{3,}/;
const PARTIAL_UUID_REGEX = /^[0-9A-F]{8}-[0-9A-F]{4}-[1-5]/i;

export const detectSearchQueryType = query => {
  if (query.match(PARTIAL_UUID_REGEX)) {
    return 'uuid';
  }

  if (query.match(PARTIAL_EMAIL_REGEX)) {
    return 'email';
  }

  if (query.match(PARTIAL_PHONE_REGEX)) {
    return 'phone';
  }

  return 'unknown';
};
