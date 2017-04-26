package org.mayhem.database;

import com.couchbase.client.java.Bucket;
import com.couchbase.client.java.Cluster;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;


@Configuration
public class Database {

   @Value("${spring.couchbase.bucket.name}")
   private String bucketName;

   @Value("${spring.couchbase.bucket.password}")
   private String bucketPassword;

   @Autowired
   private Cluster couchbaseCluster;

   public @Bean
   Bucket bucket() {
        return couchbaseCluster.openBucket(bucketName, bucketPassword);
   }
}