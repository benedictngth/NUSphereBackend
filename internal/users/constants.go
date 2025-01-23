package users

const INVALID_USERNAME_ERROR = "INVALID USERNAME"
const INVALID_USERNAME = "Invalid Username!"
const INVALID_PASSWORD_ERROR = "INVALID PASSWORD"
const INVALID_PASSWORD = "Invalid Password!"

const AUTH_USER_NOT_FOUND = "AUTH: USER NOT FOUND IN COOKIES"

const AUTH_FAILURE = "AUTH: AUTHENTICATION COOKIE NOT FOUND"
const AUTH_SUCCESS = "AUTH: AUTHENTICATION COOKIE FOUND"

const LOGOUT_SUCCESS = "AUTH: LOGOUT SUCCESSFUL; COOKIE DELETED"
const LOGOUT_FAILURE = "AUTH: LOGOUT FAILED; COOKIE NOT DELETED"

const NO_USER_PROFILE = "GET: NO USER FOUND IN PROFILE"

const REGISTRATION_FAILED = "AUTH: REGISTRATION FAILED"
const REGISTRATION_SUCCESS = "AUTH: REGISTRATION SUCCESSFUL"

const USERS_NOT_FOUND = "GET: NO USERS FOUND"

const DUPLICATE_USER_ERROR = "unable to insert row: ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)"
const DUPLICATE_USER = "Error! User exists!"

const NOT_MIN_LENGTH_ERROR = "Key: 'RegisterRequest.Password' Error:Field validation for 'Password' failed on the 'min' tag"
const NOT_MIN_LENGTH = "Password must be at least 6 characters long!"
