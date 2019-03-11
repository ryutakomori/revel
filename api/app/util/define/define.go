package define

const SUCCESS = 1

const AUTH_SUCCESS = 201

const AUTH_FAILURE = 401          //認証エラー
const AUTH_REGISTER_FAILURE = 402 //登録エラー
const AUTH_RESET_FAILURE = 403    // 存在しない
const AUTH_NOT_EXIST = 404        // 存在しない

const ERROR_VALIDATE = 422  // バリデーションエラー
const ERROR_SERVER = 500    // サーバーエラー
const ERROR_SEND_MAIL = 501 // メール送信エラー
const ERROR_EXPIRED = 502   // 期限切れ
const ERROR_UPDATE = 503    // 更新エラー

const DB_SQL_FAILURE = 900

const LIMIT_RATE = 10        //アクセス制限回数
const LIMIT_TIME_SECOND = 10 //制限時間

const BRIDGE_COMPANIES_NUM = 10         //Lambda function 連携数
const BRIDGE_COMPANIES_CREATE_NUM = 100 //Lambda function 連携数
