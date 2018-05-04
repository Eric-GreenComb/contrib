import java.security.MessageDigest;

public class Address{
	public static void main(String[] args) {	
        System.out.println(GenAddr("工作计划"));  
	}

	public static String GenAddr(String plainText){
		try {
			MessageDigest md = MessageDigest.getInstance("SHA-256");
			md.update(plainText.getBytes("UTF8"));
			md.update(md.digest());

			// for(byte b:md.digest())
			// 	System.out.format("%02x",b);	

			byte b[] = md.digest();  
  
            int i;  
  
            StringBuffer buf = new StringBuffer("");  
            for (int offset = 0; offset < b.length; offset++) {  
                i = b[offset];  
                if (i < 0)  
                    i += 256;  
                if (i < 16)  
                    buf.append("0");  
                buf.append(Integer.toHexString(i));  
            }   
            return buf.toString();  		
		} catch (Exception e) {			
			e.printStackTrace();
			return null;  
		}
	}
}