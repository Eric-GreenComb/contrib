import java.io.BufferedReader;  
import java.io.BufferedWriter;  
import java.io.FileReader;  
import java.io.FileWriter;  
import java.io.IOException;  
import java.security.InvalidKeyException;  
import java.security.KeyFactory;  
import java.security.KeyPair;  
import java.security.KeyPairGenerator;  
import java.security.NoSuchAlgorithmException;  
import java.security.SecureRandom;  
  
import java.security.interfaces.RSAPrivateKey;  
import java.security.interfaces.RSAPublicKey;  
import java.security.spec.InvalidKeySpecException;  
import java.security.spec.PKCS8EncodedKeySpec;  
import java.security.spec.X509EncodedKeySpec;  
  
import javax.crypto.BadPaddingException;  
import javax.crypto.Cipher;  
import javax.crypto.IllegalBlockSizeException;  
import javax.crypto.NoSuchPaddingException;  

import java.security.PrivateKey;  
import java.security.PublicKey;  

import java.util.Base64;

public class RSASignature {  

    /** 
     * 签名算法 
     */  
    public static final String SIGN_ALGORITHMS = "SHA256WithRSA";  

    /** 
     * 从文件中加载私钥 
     *  
     * @param keyFileName 
     *            私钥文件名 
     * @return 是否成功 
     * @throws Exception 
     */  
    public static String loadPrivateKeyByFile(String path) throws Exception {  
        try {  
            BufferedReader br = new BufferedReader(new FileReader(path  
                    + "/pkcs8_private.pem"));  
            String readLine = null;  
            StringBuilder sb = new StringBuilder();  
            while ((readLine = br.readLine()) != null) {  
                sb.append(readLine);  
            }  
            br.close();
            String rpHead = sb.toString().replace("-----BEGIN PRIVATE KEY-----","");
            return rpHead.replace("-----END PRIVATE KEY-----","");  
        } catch (IOException e) {  
            throw new Exception("私钥数据读取错误");  
        } catch (NullPointerException e) {  
            throw new Exception("私钥输入流为空");  
        }  
    } 

    /** 
     * 从文件中输入流中加载公钥 
     *  
     * @param in 
     *            公钥输入流 
     * @throws Exception 
     *             加载公钥时产生的异常 
     */  
    public static String loadPublicKeyByFile(String path) throws Exception {  
        try {  
            BufferedReader br = new BufferedReader(new FileReader(path  
                    + "/public.pem"));  
            String readLine = null;  
            StringBuilder sb = new StringBuilder();  
            while ((readLine = br.readLine()) != null) {  
                sb.append(readLine);  
            }  
            br.close();  
            String rpHead = sb.toString().replace("-----BEGIN PUBLIC KEY-----","");
            return rpHead.replace("-----END PUBLIC KEY-----","");  
        } catch (IOException e) {  
            throw new Exception("公钥数据流读取错误");  
        } catch (NullPointerException e) {  
            throw new Exception("公钥输入流为空");  
        }  
    }      

    public static String sign(String content, String privateKey)  
    {  
        try   
        {  

            PKCS8EncodedKeySpec priPKCS8    = new PKCS8EncodedKeySpec( Base64.getDecoder().decode(privateKey) );   
            KeyFactory keyf = KeyFactory.getInstance("RSA");  
            PrivateKey priKey = keyf.generatePrivate(priPKCS8); 

            java.security.Signature signature = java.security.Signature.getInstance(SIGN_ALGORITHMS);  
            signature.initSign(priKey);  
            signature.update( content.getBytes());  
            byte[] signed = signature.sign();  
            return Base64.getEncoder().encodeToString(signed);
        }  
        catch (Exception e)   
        {  
            e.printStackTrace();  
        }  
        return null;  
    } 

    public static boolean doCheck(String content, String sign, String publicKey)  
    {  
        try   
        {  
            KeyFactory keyFactory = KeyFactory.getInstance("RSA");  
            byte[] encodedKey = Base64.getDecoder().decode(publicKey);  
            PublicKey pubKey = keyFactory.generatePublic(new X509EncodedKeySpec(encodedKey));  
  
          
            java.security.Signature signature = java.security.Signature  
            .getInstance(SIGN_ALGORITHMS);  
          
            signature.initVerify(pubKey);  
            signature.update( content.getBytes() );  
          
            boolean bverify = signature.verify( Base64.getDecoder().decode(sign) );  
            return bverify;  
              
        }   
        catch (Exception e)   
        {  
            e.printStackTrace();  
        }  
          
        return false;  
    }               

    public static void main(String[] args) throws Exception {  

        String filepath="/home/eric/go/src/github.com/CebEcloudTime/tools/java/rsa";  

        String key = loadPrivateKeyByFile(filepath);

        System.out.println(key);
  
        //RSAEncrypt.genKeyPair(filepath);  
          
        System.out.println("---------------私钥签名过程------------------");  
        String content = "用于签名的原始数据";  

        // key = "MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJ0DWxqG4kkQYPoJMbvmpy+vjsgQ17DemzIcGnU9OYwPmhfmSywVKlmrFVyGgFlOjSHJK5+x3oqkET5YbTOviGjoy/uV4Ec0jCCjS3IFc59vMHV5a3yir8o3dwInmnuNBKnaPFTl10dreeAhtcmDF7hZ/tYnnG9KdrODOSDE/pTRAgMBAAECgYBt+9LOQxuxWHLF0rjuyUPlSFF43StpbpVBxaPW6fssnCUxhpSznWPcCdZdyK2RYU/FEdin9X1Qmlql1GUyJkwjL6v+vMlN9o4AkYqzOVJnu6vHRgLB/BP78C1b7MHn+l+SPyLApBe/daeY9XZfOlD/S4Enjrvg4DsDfXnkouYWAQJBAM37rGhzn9IRi/C9wNjhLXwuRz6ubD9dF2TC1lIuZDXzFVt9f9PXxvsrwKWsTayK/SkN1MIA22q9SmmaAQvXP/ECQQDDI5vPkPa6UHjZ3FVEkG6DgmD847WJZNgmtNb/dARPwqnmOQrYZAarDZ3E+Fq/uVFDTtGEjJtfqpQhkk76x4LhAkAE4yrVE6FAJ8BtRuNTggxFPQfdud/BpSDP+DuDmawxB4KDODgXO7Bx9zjL9YmmRWn6VmSs8b5DCxi/5rKNqF7RAkBZw5CR+8ozTH87IGqs3o+nuRrqWckRSa1QqNFZs0GkexRyjfzaK7ERkHLpv6DnHtUt1Bz3D0MNz8bSZp4kKBChAkBq/2Lwf1Q35280F6Hd9t98WY2XKIvXjskvNfDlxRkV0L/U6fvKp5S2fRXmxDpMvShYqmoAbC3WYz3C6Tp1kF7r";
        String signstr = sign(content,key);  
        System.out.println("签名原串："+content);  
        System.out.println("签名串：");
        System.out.println(signstr);  
        System.out.println();  
          
        System.out.println("---------------公钥校验签名------------------");  
        System.out.println("验签结果："+RSASignature.doCheck(content, signstr, loadPublicKeyByFile(filepath)));  
        System.out.println();  
          
    }  
} 