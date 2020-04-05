function myFunction() {
  var props = PropertiesService.getScriptProperties();
  var accessKey = props.getProperty('S3_ACCESS_KEY_ID');
  var secretKey = props.getProperty('S3_SECRET_ACCESS_KEY');
  var bucketName = props.getProperty('BUCKET_NAME');
  var filePath = props.getProperty('UPLOAD_FILE_PATH');
  var fileId = props.getProperty('FILE_ID');

  // スプレッドシートを取得
  var ss = SpreadsheetApp.openById(fileId);

  // シートのオブジェクトを取得
  var sheet = ss.getSheets()[0];

  // データを取得
  var data = sheet.getRange('A:C').getValues();

  // 送信データ用の配列を用意
  var csv = '';

  // データをチェックしながらループ
  for ( var i = 0; i < data.length; i++ )
  {
    // id があれば
    if ( data[i][0] != '' )
    {
      // データを作成
      csv += '"' + data[i][0] + '","' + data[i][1] + '","' + data[i][2] + '"' + "\n";
    }
  }

  // バイナリに変換
  csv = Utilities.newBlob( csv );
  
  var s3 = S3.getInstance( accessKey, secretKey );
  s3.putObject( bucketName, filePath, csv, {logRequests:true} );
}
