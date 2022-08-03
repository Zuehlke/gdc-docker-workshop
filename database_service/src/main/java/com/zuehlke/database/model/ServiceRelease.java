package com.zuehlke.database.model;

import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.time.Instant;

@Entity
public class ServiceRelease {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    public Long id;

    @Column(unique = true, nullable = false)
    public String releaseName;

    @Column(unique = true, nullable = false)
    public String releaseTitle;

    @Column(nullable = false)
    @CreationTimestamp
    public Instant releaseDate;

}
