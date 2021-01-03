export const STATUS_DESCRIPTIONS = {
  UNCONFIRMED: 'User has been created but not confirmed',
  CONFIRMED: 'User has been confirmed',
  ARCHIVED: 'User is no longer active',
  COMPROMISED: 'User is disabled due to a potential security threat',
  UNKNOWN: 'User status is not known',
  RESET_REQUIRED: 'User is confirmed, but the user must request a code and reset his or her password before he or she can sign in',
  FORCE_CHANGE_PASSWORD: 'The user is confirmed and the user can sign in using a temporary password, but on first sign-in, the user must change his or her password to a new value before doing anything else',
};
