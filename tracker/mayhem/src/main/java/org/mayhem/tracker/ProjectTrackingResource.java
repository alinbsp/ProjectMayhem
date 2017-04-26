package org.mayhem.tracker;

import com.couchbase.client.java.Bucket;
import com.couchbase.client.java.document.JsonDocument;
import com.couchbase.client.java.document.json.JsonObject;
import com.couchbase.client.java.query.N1qlQuery;
import com.couchbase.client.java.query.N1qlQueryResult;
import com.couchbase.client.java.query.Statement;
import org.mayhem.tracker.domain.AddressesResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.dao.DataRetrievalFailureException;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import java.util.UUID;
import java.util.stream.Collectors;

import static com.couchbase.client.java.query.Select.select;
import static com.couchbase.client.java.query.dsl.Expression.i;
import static com.couchbase.client.java.query.dsl.Expression.s;
import static com.couchbase.client.java.query.dsl.Expression.x;

@RestController
@RequestMapping("/track")
public class ProjectTrackingResource {

    private final Bucket bucket;

    @Value("${storage.expiry:0}")
    private int expiry;

    @Autowired
    public ProjectTrackingResource(@Qualifier("bucket") Bucket bucket) {
        this.bucket = bucket;
    }

    @RequestMapping(value = "/register/{projectId}", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON)
    public ResponseEntity register(@Context HttpServletRequest httpServletRequest, @PathVariable("projectId") final String projectId) {
        final String referer = httpServletRequest.getRemoteAddr();
        if(checkProjectAndAddress(referer, projectId)) {
            JsonObject data = JsonObject.create()
                    .put("project", projectId)
                    .put("address", referer);
            JsonDocument doc;
            if (expiry > 0) {
                doc = JsonDocument.create("entry::" + UUID.randomUUID().toString(), expiry, data);
            } else {
                doc = JsonDocument.create("entry::" + UUID.randomUUID().toString(), data);
            }

            try {
                bucket.insert(doc);
            } catch (Exception e) {
                throw new RuntimeException("There was an error adding entry to project tracker");
            }
        }

        return ResponseEntity.ok(getAllReferersForProject(projectId));
    }

    private boolean checkProjectAndAddress(final String referer, final String projectId) {
        Statement query = select(x("project"))
                .from(i(bucket.name()))
                .where(x("project").eq(s(projectId)).and(x("address").eq(s(referer))));

        N1qlQueryResult result = bucket.query(N1qlQuery.simple(query));

        if (!result.finalSuccess()) {
            throw new DataRetrievalFailureException("Query error: " + result.errors());
        }

        return result.allRows().size() == 0;
    }

    private AddressesResponse getAllReferersForProject(final String projectId) {
        Statement query = select(x("address"))
                .from(i(bucket.name()))
                .where(x("project").eq(s(projectId)));

        N1qlQueryResult result = bucket.query(N1qlQuery.simple(query));

        if (!result.finalSuccess()) {
            throw new DataRetrievalFailureException("Query error: " + result.errors());
        }

        return new AddressesResponse(result.allRows().stream().map(row -> row.value().getString("address")).collect(Collectors.toList()));
    }
}
