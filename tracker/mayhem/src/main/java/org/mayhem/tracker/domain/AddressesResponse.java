package org.mayhem.tracker.domain;

import java.util.List;

public class AddressesResponse {
    private List<String> addresses;

    public AddressesResponse() {
    }

    public AddressesResponse(List<String> addresses) {
        this.addresses = addresses;
    }

    public List<String> getAddresses() {
        return addresses;
    }

    public void setAddresses(List<String> addresses) {
        this.addresses = addresses;
    }
}
