package com.zuehlke.database.controller;

import com.zuehlke.database.model.ServiceRelease;
import com.zuehlke.database.model.ServiceReleaseRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.util.List;

@RestController
@RequiredArgsConstructor
@RequestMapping("/releases")
public class ServiceReleaseController {

    private final ServiceReleaseRepository serviceReleaseRepository;

    @PostMapping
    public ServiceRelease scheduleRelease(@Valid @RequestBody ServiceRelease serviceRelease) {
        return serviceReleaseRepository.save(serviceRelease);
    }

    @GetMapping
    public List<ServiceRelease> getAllReleases() {
        return serviceReleaseRepository.findAll();
    }

}
