package com.zuehlke.database.model;


import javax.persistence.*;
import javax.validation.constraints.Future;
import javax.validation.constraints.NotBlank;
import java.time.Instant;

@Entity
public class ServiceRelease {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    public Long id;

    @Column(unique = true, nullable = false)
    @NotBlank(message = "Release title must not be blank!")
    public String releaseTitle;

    @Column(unique = true, nullable = false)
    @NotBlank(message = "Release version must not be blank!")
    public String releaseVersion;

    @Column(nullable = false)
    @Future(message = "How do you plan to release in the past?")
    public Instant releaseDate;

}
