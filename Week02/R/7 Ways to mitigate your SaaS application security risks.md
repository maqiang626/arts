# 			7 Ways to mitigate your SaaS application security risks		

[CNCF 				Member				Blog Post](https://www.cncf.io/lf-author-category/member/)

Posted on 			October 8, 2021			 					 					

*Originally published on the [Msys Technologies blog](https://www.msystechnologies.com/blog/7-ways-to-mitigate-your-saas-application-security-risks/)*

If you’re a SaaS entrepreneur or you’re looking to build a SaaS  application, in that case, you may already be aware of the fact that  there is a new economy that has evolved around SaaS (Software as a  Service). Core business services are offered to the consumers as a  subscription model via pay-per-use in this SaaS market. Studies have  revealed that Software as a service (SaaS) enterprises are evolving at a sky-rocket speed. They are becoming the first choice due to their  simple up-gradation, scalability, and low infrastructure obligations.  Per [Smartkarrot.com](https://www.smartkarrot.com/resources/blog/saas-statistics/), the SaaS industry’s market capitalization in 2020 was approximately  $110 Billion and is expected to touch the $126 billion mark by the end  of 2021. And it is expected to reach $143 billion by the year 2022.

However, security is one of the primary reasons why small and medium  businesses hold back from taking full advantage of powerful cloud  technologies. Though the total cost of ownership was once viewed as the  main blockage for possible SaaS customers, security is now on top of  that list. The anxieties with SaaS security evolved with more and more  users embracing the new technology, but is everything all that bad as  reviews and opinions hint? Here are 7 SaaS security best practices that  can help you in curbing SaaS security risks, that too cost-effectively:

### 1. Use a Powerful Hosting Service (AWS, Azure, GCP, etc.) and Make Full Use of their Security

The biggest cloud providers have spent millions of dollars on  security research and development and made it available worldwide.  Leverage their infrastructure and the best SaaS cybersecurity practices  that they have made available to the public and focus your energy on the core issue(s) your software resolves.

**a. API Gateway Services**

**b. Security Monitoring Services**

**c. Encryption Services**

### 2. SaaS Application Security — Reduce Attack Surface and Vectors

**a. Software/Hardware –** For example, do not define  endpoints in your public API for admin related tasks. If the endpoint  doesn’t exist, there is nothing else to secure (when it comes to SaaS  endpoint protection)!

**b. People –** Limit the access people have to any  sensitive data. If required, for a user to access sensitive data, log  all the actions taken and, if possible, make it necessary to have more  than one person involved in accessing the data.

### 3. SaaS Security Checklist — Do not Save Sensitive Data

**a.** Only capture data you absolutely need. For  instance, if you never use a person’s national ID number (e.g., SSN),  don’t ask for it)

**b.** Assign a third party for sensitive data storing.

In this, for example, your system never holds possession of any  credit card number, so you don’t have to worry about protecting it.

### 4. Encrypt all your Customer Data — Adopt the Best SaaS Security Solutions

**a. Data at Rest:** When any data is saved either as a  file or inside a database, it is considered “at rest.” Almost every data storage service can store the data when it is encrypted and then  decrypt it when you ask for it. For example, SQL Server enables you to  turn on a setting to encrypt the stored data with their Transparent Data Encryption (TDE) feature.

**b. Data in Flight:** When data is read from storage  and transferred out of the currently running process, it is called  “in-flight.” Sending data over any networking protocol, be it FTP, TCP,  HTTP, is data that is “in-flight.” Network sniffers (if attached to your network) can read this data, and if it is not encrypted, it can be  stolen. Employing SSL/TLS for HTTP is a typical example.

### 5. Log All Access and Modifications to Sensitive Data — Opt for a Robust SaaS Security Architecture

There’s no guarantee that your system’s security will never be  breached. It is more of a question of “when will it happen” rather than  “if it will happen.” For this very reason, it is crucial to log all  changes and access to stored sensitive data and adjustments to user  permissions and login attempts. When something actually goes wrong, you  have an audit log that can be used to solve how the breach occurred and  know what needs to change to stop any further similar security breaches.

### 6. Implement Two-factor Authentication

Social engineering is the most common way which hackers use to breach any system. Make social engineering hacks more complex by asking users  to have a second step to authenticate with your system. Implement a  system that needs at least two of the following three types of  information:

- Something the user knows (e.g., username/password)
- Something the user has (e.g., phone)
- Something the user is (e.g., fingerprint)

Sending a code to a user’s phone or email is a simple yet effective  way to implement two-factor authentication. To balance the added  security with the demand for usability, give your clients the option of  choosing if they would like to use the phone or email and an option for  the code validity for the device being used.

### 7. Use a Key Vault Service

Key Vaults allow the stored sensitive data to be accessed only by  applications that have been given access to the Key Vault, removing the  need for a person to handle the secrets. A Key Vault stores all secrets  to encrypt data, databases/datastores access, electronically signed  files, etc. Cloud platforms like Azure and AWS offer highly secure and  configurable Key Vault services.

For extra security, use different key vaults for different customers. For advanced security, allow your customers to bring their keys.

### Takeaway

There are several reasons why businesses must take advantage of cloud computing to enhance their operational efficiency and reduce their  costs. Nevertheless, security concerns often hold back businesses from  placing their valuable data in the cloud. But, with the right technology and best practices, SaaS can be far more secure than any on premise  application, and you can have numerous options for retaining control  over your security infrastructure and address the security issues  head-on with your respective provider.





------

安全：

安全问题已经提升日程，所有软件从业者都不能忽视



这里有 7 个 SaaS 安全最佳实践，可以帮助您遏制 SaaS 安全风险，而且成本效益很高：

1. 使用强大的托管服务（AWS、Azure、GCP 等）并充分利用它们的安全性
2. SaaS 应用安全——减少攻击面和向量
3. SaaS 安全检查表——不要保存敏感数据
4. 加密您的所有客户数据——采用最佳 SaaS 安全解决方案
5. 记录对敏感数据的所有访问和修改——选择强大的 SaaS 安全架构
6. 实现两因素身份验证
7. 使用 Key Vault 服务



结合目前对安全知识的学习，结合公司的实际应用，不断总结！

