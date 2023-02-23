import URLFetchRequestOptions = GoogleAppsScript.URL_Fetch.URLFetchRequestOptions;

const doPost = () => {
  // リクエストヘッダ
  const headers = {
    "Content-Type": "application/json; charset=UTF-8",
    "Authorization": `Bearer ${LINE_ACCESS_TOKEN}`,
  };

  // メッセージ
  const postData = {
    to: LINE_USER_ID,
    messages: [
      {
        type: "text",
        text: `今日の使用した分の金額をZaimに入力しましたか？ `,
      },
    ],
  };

  // POSTオプション作成
  const options: URLFetchRequestOptions = {
    method: "post",
    headers: headers,
    payload: JSON.stringify(postData),
  };

  return UrlFetchApp.fetch(PUSH_API, options);
};
